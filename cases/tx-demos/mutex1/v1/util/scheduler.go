package util

import (
	"context"
	"fmt"
	"playground/cases/model"
	"playground/cases/tx-demos/mutex1/util"
	"sync"
	"time"

	"code.guanmai.cn/back_end/grpckit"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	"github.com/google/uuid"

	grpckitsync "code.guanmai.cn/back_end/grpckit/sync"
	"gorm.io/gorm"
)

const (
	DistributedMutexExpiration = 10 * time.Second // 分布式锁超时时间

	SchedulerTickerInterval   = 5 * time.Second        // 调度器默认触发间隔
	SchedulerMainLoopInterval = 100 * time.Millisecond // 调度器休眠时间间隔

	TaskPrepareTimeout     = 5 * time.Minute  // 任务PrepareTask调用超时时间
	TaskExecuteTimeout     = 2 * time.Minute  // 任务ExecuteTask调用超时时间
	TaskFailureInterval    = 2 * time.Second  // 任务失败内置调度间隔
	TaskScheduleExpiration = 10 * time.Minute // 任务调度状态过期时间
)

type Scheduler struct {
	TaskDelegate     *model.TaskDelegate
	Mutex            sync.Mutex
	DistributedMutex grpckitsync.Mutex
	ConcurrentTask   uint32 // 当前并发任务数
	TaskClearTime    uint64 // 上次清理调度状态时间
	IsRunning        bool
}

func NewScheduler(delegate *model.TaskDelegate) (*Scheduler, error) {
	s := &Scheduler{
		TaskDelegate:     delegate,
		DistributedMutex: impl.Sync.NewMutex(fmt.Sprintf("/scheduler/%d", delegate.Type), DistributedMutexExpiration),
	}

	return s, nil
}

func (s *Scheduler) Start() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if !s.IsRunning {
		s.IsRunning = true
		grpclog.Infof("scheduler %d start %s", s.TaskDelegate.Type, s.TaskDelegate.Endpoint)

		go func() {
			s.MainLoop()
		}()
	}
}

func (s *Scheduler) Stop() {
	// TODO
}

func (s *Scheduler) MainLoop() {
	ticker := time.NewTicker(SchedulerTickerInterval)
	done := make(chan bool, 128)

	for {
		select {
		case <-ticker.C:
			//grpclog.Infof("scheduler %d ticker", s.TaskDelegate.Type)
		case <-done:
			//grpclog.Infof("scheduler %d done", s.TaskDelegate.Type)
		}

		tasks, err := s.GetNextTasks()

		if err != nil {
			grpclog.Infof("scheduler %d next tasks error: %s", s.TaskDelegate.Type, err)
		} else if len(tasks) == 0 {
			grpclog.Infof("scheduler %d noop", s.TaskDelegate.Type)
		} else {
			for _, t := range tasks {
				go func(task *model.Task) {
					_, _ = s.Schedule(task)
					done <- true
				}(t)
			}
		}

		now := util.UnixMilliNow()
		min := now - uint64(TaskScheduleExpiration.Milliseconds())
		if s.TaskClearTime < min {
			// 清理过期的任务调度状态
			sql := fmt.Sprintf("UPDATE %s set is_busy = false WHERE type = %d AND is_busy = true AND last_schedule_time < %d", GetTaskTable(), s.TaskDelegate.Type, min)
			s.Impl.DB.Exec(sql)
			s.TaskClearTime = now
		}

		// Sleep下避免CPU爆炸
		time.Sleep(SchedulerMainLoopInterval)
	}
}

func (s *Scheduler) GetClient() (asynctaskapi.AsyncTaskDelegateServiceClient, error) {
	if s.Client == nil {
		s.Mutex.Lock()
		defer s.Mutex.Unlock()

		if s.Client == nil {
			conn, err := grpckit.Dial(
				s.TaskDelegate.Endpoint,
			)
			if err != nil {
				return nil, err
			}

			s.Client = asynctaskapi.NewAsyncTaskDelegateServiceClient(conn)
			grpclog.Infof("scheduler %d connect %s", s.TaskDelegate.Type, s.TaskDelegate.Endpoint)
		}
	}
	return s.Client, nil
}

func (s *Scheduler) Schedule(task *model.Task) (*model.Task, error) {
	client, err := s.GetClient()
	if err != nil {
		return nil, err
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	requestId := u.String()

	grpclog.Infof("scheduler %d task %d start request id %s", s.TaskDelegate.Type, task.TaskId, requestId)

	ctx := context.Background()
	if task.UserInfo != nil {
		creatorCtx, err := commonauth.NewOutgoingContext(task.UserInfo)
		if err != nil {
			return nil, err
		}
		ctx = metadata.AppendToOutgoingContext(creatorCtx, "x-request-id", requestId)
	}

	defer func() {
		s.Mutex.Lock()
		s.ConcurrentTask--
		s.Mutex.Unlock()

		task.IsBusy = false
		if task.State != model.Task_STATE_COMPLETED && task.State != model.Task_STATE_CANCELED && task.State != model.Task_STATE_FAULTED {
			if task.FailureCount >= s.TaskDelegate.TaskMaxFailureCount || task.ScheduleCount >= s.TaskDelegate.TaskMaxSchduleCount {
				task.State = model.Task_STATE_FAULTED
			}
		}

		_, _ = s.UpdateTask(task)
		grpclog.Infof(
			"scheduler %d task %d update state %d schedule %d/%d failure %d/%d",
			s.TaskDelegate.Type,
			task.TaskId, task.State, task.ScheduleCount, s.TaskDelegate.TaskMaxSchduleCount, task.FailureCount, s.TaskDelegate.TaskMaxFailureCount,
		)
	}()

	if task.State == model.Task_STATE_CREATED {
		ctx, cancel := context.WithTimeout(ctx, TaskPrepareTimeout)
		defer cancel()

		processTaskResponse, err := client.PrepareTask(ctx, &asynctaskapi.PrepareTaskRequest{
			Task: task,
		})
		if err != nil {
			grpclog.Infof("scheduler %d task %d error: %s", s.TaskDelegate.Type, task.TaskId, err)
			task.FailureCount++
			task.NextScheduleTime = task.LastScheduleTime + uint64(TaskFailureInterval/time.Millisecond)
			return nil, err
		}

		if processTaskResponse.TaskState != task.State && processTaskResponse.TaskState != model.Task_STATE_UNSPECIFIED {
			task.State = processTaskResponse.TaskState
		} else {
			task.State = model.Task_STATE_READY
		}
		if processTaskResponse.TaskData != nil {
			taskData := processTaskResponse.TaskData
			taskData.TaskDataId = task.TaskId
			_, _ = s.UpdateTaskData(taskData)
		}

	} else if task.State == model.Task_STATE_READY || task.State == model.Task_STATE_RUNNING {
		ctx, cancel := context.WithTimeout(ctx, TaskExecuteTimeout)
		defer cancel()

		taskData, err := s.GetTaskData(task.TaskId)
		if err != nil {
			grpclog.Infof("scheduler %d task %d error: %s", s.TaskDelegate.Type, task.TaskId, err)
			return nil, err
		}

		executeTaskResponse, err := client.ExecuteTask(ctx, &asynctaskapi.ExecuteTaskRequest{
			Task:     task,
			TaskData: taskData,
		})
		if err != nil {
			grpclog.Infof("scheduler %d task %d error: %s", s.TaskDelegate.Type, task.TaskId, err)
			task.FailureCount++
			task.NextScheduleTime = task.LastScheduleTime + uint64(TaskFailureInterval/time.Millisecond)
			return nil, err
		}

		if executeTaskResponse.TaskState != task.State && executeTaskResponse.TaskState != model.Task_STATE_UNSPECIFIED {
			task.State = executeTaskResponse.TaskState
		}
		if executeTaskResponse.TaskData != nil {
			taskData := executeTaskResponse.TaskData
			taskData.TaskDataId = task.TaskId
			_, _ = s.UpdateTaskData(taskData)
		}
	}

	return task, nil
}

func (s *Scheduler) GetNextTasks() ([]*model.Task, error) {
	err := s.DistributedMutex.Lock()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = s.DistributedMutex.Unlock()
	}()
	amount := s.TaskDelegate.MaxConcurrentTask - s.ConcurrentTask
	if amount == 0 {
		return nil, nil
	}

	var tasks []*model.Task
	err = s.Impl.DB.Transaction(func(tx *gorm.DB) error {
		now := commonutil.UnixMilliNow()
		result := tx.Where(
			fmt.Sprintf("type = ? AND state IN ? AND is_busy = false AND next_schedule_time < ? AND singleton_id not in (SELECT singleton_id FROM %s WHERE type = %d AND is_busy = true)", GetTaskTable(), s.TaskDelegate.Type),
			s.TaskDelegate.Type,
			[]int32{
				int32(model.Task_STATE_CREATED),
				int32(model.Task_STATE_READY),
				int32(model.Task_STATE_RUNNING),
			},
			now,
		).Order("priority").Order("last_schedule_time").Limit(int(amount)).Find(&tasks)
		if result.Error != nil {
			return result.Error
		}

		now = commonutil.UnixMilliNow()
		for _, task := range tasks {
			task.IsBusy = true
			task.ScheduleCount++
			task.LastScheduleTime = now
			result := tx.Save(task)
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	s.ConcurrentTask += uint32(len(tasks))

	return tasks, nil
}

func (s *Scheduler) UpdateTask(task *model.Task) (*model.Task, error) {
	handler := s.Impl.DB

	result := handler.Save(task)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (s *Scheduler) GetTaskData(taskId uint64) (*model.TaskData, error) {
	handler := s.Impl.DB

	taskData := &model.TaskData{}

	result := handler.First(taskData, taskId)
	if result.Error != nil {
		return nil, result.Error
	}

	return taskData, nil
}

func (s *Scheduler) UpdateTaskData(taskData *model.TaskData) (*model.TaskData, error) {
	handler := s.Impl.DB

	result := handler.Save(taskData)
	if result.Error != nil {
		return nil, result.Error
	}

	return taskData, nil
}
