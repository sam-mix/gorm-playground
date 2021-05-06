主机 
gmadmin@A-task-worker-03


cd /data/www/station/hotfix/sorter_perf_cal/gm_service
rm -rf venv

wget https://bootstrap.pypa.io/3.5/get-pip.py
python3 get-pip.py


mkdir sorter_perf_cal
cd sorter_perf_cal
git clone git@code.guanmai.cn:back_end/gm_service.git
cd gm_service
git checkout hotfix/sorter_perf_cal
python3 -m venv venv
source ./venv/bin/activate
pip install --upgrade "pip < 21.0"
pip install -r requirements.txt
vim /data/www/station/hotfix/sorter_perf_cal/gm_service/tools/crontab_tasks/cal_sorter_perf.py
python3 /data/www/station/hotfix/sorter_perf_cal/gm_service/tools/crontab_tasks/cal_sorter_perf.py


