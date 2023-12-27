package category_brand_service

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

type DeleteBrandByCategoryService struct {
	ctx context.Context
}

// NewDeleteBrandByCategoryService new DeleteBrandByCategoryService
func NewDeleteBrandByCategoryService(ctx context.Context) *DeleteBrandByCategoryService {
	return &DeleteBrandByCategoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteBrandByCategoryService) Run(req *common.Req) (resp *common.Empty, err error) {
	// Delete all records with the category id
	res := db.GetMysql().Where("category_id = ?", req.Id).Delete(&model.CategoryBrand{})
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("the record to delete does not exist")
		return nil, kerrors.NewBizStatusError(code.NotFound, "Deletion failed. The related categories or brands are not existed.")
	}

	return &common.Empty{}, nil
}
