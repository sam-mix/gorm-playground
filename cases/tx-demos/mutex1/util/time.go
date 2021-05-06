package util

import "time"

const NanosecondsPerMillisecond = uint64(time.Millisecond) / uint64(time.Nanosecond)

// UnixMilliNow 返回当前的 Unix 时间，使用 Millisecond 作为单位
func UnixMilliNow() uint64 {
	var now int64 = time.Now().UnixNano() / int64(NanosecondsPerMillisecond)
	return uint64(now)
}
