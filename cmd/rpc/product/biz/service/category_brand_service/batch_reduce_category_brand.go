package category_brand_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type BatchReduceCategoryBrandService struct {
	ctx context.Context
}

// NewBatchReduceCategoryBrandService new BatchReduceCategoryBrandService
func NewBatchReduceCategoryBrandService(ctx context.Context) *BatchReduceCategoryBrandService {
	return &BatchReduceCategoryBrandService{ctx: ctx}
}

// Run create note info
func (s *BatchReduceCategoryBrandService) Run(req *product.NewCategoryBrand_) (resp *common.Empty, err error) {
	res := db.GetMysql().Where("brand_id = ?", req.BrandId).
		Where("category_id in ?", req.CategoryId).Delete(&model.CategoryBrand{})
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("the record to delete does not exist")
		return nil, kerrors.NewBizStatusError(code.NotFound, "Deletion failed. The related categories or brands are not existed.")
	}

	return &common.Empty{}, nil
}
