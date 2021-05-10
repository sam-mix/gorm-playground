kubectl get pods -n env-hotfix-aftersale 
kubectl get pods -n env-hotfix-aftersale | grep order
kubectl get pods -n env-hotfix-aftersale 

kubectl -n env-hotfix-aftersale logs -l app=order -- grep "ff5a0bde-79cd-40e9-ab1d-c756a4644219"
kubectl -n env-hotfix-aftersale logs -l app=order -- grep "key => "
kubectl -n env-hotfix-aftersale logs -l app=ceres-aftersale -- grep "key => "

kubectl -n env-hotfix-aftersale logs -l app=ceres-after-sale -f | grep "key => "

kubectl -n env-hotfix-aftersale logs -l app=ceres-order -f | grep "4b9f2099-b576-4062-a6ba-975dfbac5739"

ceres-order-6d6446ddcb-j7rn6


346179075888381994_5



curl -d '{"more":false}' "https://env-hotfix-aftersale.x.k8s.guanmai.cn/ceres/initialization/InitializationService/InitializeTestingGroup"


kubectl -n env-feature-bshop-qr-sign-in logs -l app=ceres-order -f | grep "777d9b35-d60c-47dc-94a0-2804c4b48d1e"


https://env-feature-sales-data.x.k8s.guanmai.cn/erp#/home


curl -d '{"more":false}' "https://env-feature-sales-data.x.k8s.guanmai.cn/ceres/initialization/InitializationService/InitializeTestingGroup"


kubectl -n env-feature-sales-data logs -l app=ceres-order -f | grep "reqCreateOrderDetails"
kubectl -n env-feature-sales-data logs -l app=ceres-order -f | grep "reqUpdateOrderDetails"

curl -d '{"more":false}' "https://env-feature-sales-data.x.k8s.guanmai.cn/ceres/initialization/InitializationService/InitializeTestingGroup"
{
 "group_id": "353583452303392792",
 "admin_group_user_username": "admin",
 "admin_group_user_phone": "12312341234",
 "admin_group_user_password": "123456",
 "driver_group_user_username": "driver",
 "driver_group_user_password": "123456",
 "purchase1_group_user_username": "purchaser1",
 "purchase1_group_user_password": "123456",
 "purchase2_group_user_username": "purchaser2",
 "purchase2_group_user_password": "123456",
 "customer_user_username": "customer_user1",
 "customer_user_phone": "12312341001",
 "customer_user_password": "123456"
}


分支: optimize/gorm_v2
namespace: env-optimize-gorm-v2 

curl -d '{"more":false}' "https://env-optimize-gorm-v2.x.k8s.guanmai.cn/ceres/initialization/InitializationService/InitializeTestingGroup"


kubectl get service -n env-optimize-gorm-v2 



curl -d '{"more":false}' "https://env-hotfix-aftersale.x.k8s.guanmai.cn/ceres/initialization/InitializationService/InitializeTestingGroup"

{
 "group_id": "346179071408865304",
 "admin_group_user_username": "admin",
 "admin_group_user_phone": "12312341234",
 "admin_group_user_password": "123456",
 "driver_group_user_username": "driver",
 "driver_group_user_password": "123456",
 "purchase1_group_user_username": "purchaser1",
 "purchase1_group_user_password": "123456",
 "purchase2_group_user_username": "purchaser2",
 "purchase2_group_user_password": "123456",
 "customer_user_username": "customer_user1",
 "customer_user_phone": "12312341001",
 "customer_user_password": "123456"
}


curl -d '{"more":false}' "https://env-optimize-gorm-v2.x.k8s.guanmai.cn/ceres/initialization/InitializationService/InitializeTestingGroup"
{
 "group_id": "354155366658867224",
 "admin_group_user_username": "admin",
 "admin_group_user_phone": "12312341234",
 "admin_group_user_password": "123456",
 "driver_group_user_username": "driver",
 "driver_group_user_password": "123456",
 "purchase1_group_user_username": "purchaser1",
 "purchase1_group_user_password": "123456",
 "purchase2_group_user_username": "purchaser2",
 "purchase2_group_user_password": "123456",
 "customer_user_username": "customer_user1",
 "customer_user_phone": "12312341001",
 "customer_user_password": "123456"
}


kubectl get po -n env-optimize-gorm-v2 | grep -v R



