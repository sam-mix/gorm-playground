grpckitcli create-service --name foo --port 6666 # 创建服务




修改前后:
!gorm.IsRecordNotFoundError(err)
!errors.Is(err, gorm.ErrRecordNotFound)


paging.Offset
int(paging.Offset)

paging.Limit + 1
int(paging.Limit + 1)

syncOptions.Limit
int(syncOptions.Limit)


.Count(&res)
count := int64(res) .Count(&count)

// res uint32
// res int64


.Update(
.Updates(







"context"

"erros"

// "code.guanmai.cn/back_end/ceres/common/auth"
grpcerrors "code.guanmai.cn/back_end/ceres/common/errors"
commonproto "code.guanmai.cn/back_end/ceres/common/proto"
"code.guanmai.cn/back_end/ceres/common/util"
"code.guanmai.cn/back_end/ceres/foo/proto"
idgeneratorproto "code.guanmai.cn/back_end/ceres/idgenerator/proto"
"%s"
"google.golang.org/grpc/grpclog"
"gorm.io/gorm"




	_ = context.Canceled
	_ = idgeneratorproto.NewIDGeneratorServiceClient
	_ = gorm.Expr
	// _ = auth.GetUserInfo
	_ = grpcerrors.GRPCError
	_ = commonproto.Image_TYPE_UNSPECIFIED
	_ = util.GetIdSn
	_ = idgeneratorproto.NewIDGeneratorServiceClient
	_ = proto.FooDetail{}
	_ = errors.As
	_ = grpclog.Error
