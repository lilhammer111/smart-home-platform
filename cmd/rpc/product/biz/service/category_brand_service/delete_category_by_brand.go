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

type DeleteCategoryByBrandService struct {
	ctx context.Context
}

// NewDeleteCategoryByBrandService new DeleteCategoryByBrandService
func NewDeleteCategoryByBrandService(ctx context.Context) *DeleteCategoryByBrandService {
	return &DeleteCategoryByBrandService{ctx: ctx}
}

// Run create note info
func (s *DeleteCategoryByBrandService) Run(req *common.Req) (resp *common.Empty, err error) {
	// Delete all records with the brand id
	res := db.GetMysql().Where("brand_id = ?", req.Id).Delete(&model.CategoryBrand{})
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
