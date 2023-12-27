package category_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteCategoryService struct {
	ctx context.Context
}

// NewDeleteCategoryService new DeleteCategoryService
func NewDeleteCategoryService(ctx context.Context) *DeleteCategoryService {
	return &DeleteCategoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteCategoryService) Run(req *common.Req) (resp *common.Empty, err error) {
	res := db.GetMysql().Delete(&model.Category{}, req.Id)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Infof("the category of id %d does not exist", req.Id)
		return nil, kerrors.NewBizStatusError(code.NotFound, "Deletion failed. The category is not existed.")
	}

	return &common.Empty{}, nil
}
