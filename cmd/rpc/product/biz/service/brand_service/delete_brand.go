package brand_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteBrandService struct {
	ctx context.Context
}

// NewDeleteBrandService new DeleteBrandService
func NewDeleteBrandService(ctx context.Context) *DeleteBrandService {
	return &DeleteBrandService{ctx: ctx}
}

// Run create note info
func (s *DeleteBrandService) Run(req *common.Req) (resp *common.Empty, err error) {
	res := db.GetMysql().Delete(&model.Brand{}, req.Id)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("the record to delete does not exist")
		return nil, kerrors.NewBizStatusError(code.NotFound, "The brand is not existed.")
	}

	return &common.Empty{}, nil
}
