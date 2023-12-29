package category_brand_service

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type UpdateCategoryBrandService struct {
	ctx context.Context
}

// NewUpdateCategoryBrandService new UpdateCategoryBrandService
func NewUpdateCategoryBrandService(ctx context.Context) *UpdateCategoryBrandService {
	return &UpdateCategoryBrandService{ctx: ctx}
}

// Run create note info
func (s *UpdateCategoryBrandService) Run(req *product.NewCategoryBrand_) (resp *common.Empty, err error) {
	// Check if the brand id is existed
	err = db.GetMysql().First(&model.Brand{}, req.BrandId).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The brand is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	// Check if the category id in list is existed
	err = db.GetMysql().Where("id in ?", req.CategoryId).Count().Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The brand is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	// Batch create the record and omit the conflict error
}
