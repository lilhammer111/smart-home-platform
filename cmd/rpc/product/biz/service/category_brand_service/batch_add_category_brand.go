package category_brand_service

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BatchAddCategoryBrandService struct {
	ctx context.Context
}

// NewBatchAddCategoryBrandService new BatchAddCategoryBrandService
func NewBatchAddCategoryBrandService(ctx context.Context) *BatchAddCategoryBrandService {
	return &BatchAddCategoryBrandService{ctx: ctx}
}

// Run create note info
func (s *BatchAddCategoryBrandService) Run(req *product.NewCategoryBrand_) (resp []*product.CategoryBrandInfo, err error) {
	err = db.GetMysql().First(&model.Brand{}, req.BrandId).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The brand is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	// todo test_0001
	categoryInfo := make([]model.Category, 0)
	res := db.GetMysql().Find(&categoryInfo, req.CategoryId)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no categories.")
		return nil, kerrors.NewBizStatusError(code.NotFound, "None of categories exist.")
	}

	newCategoryBrands := make([]model.CategoryBrand, 0)
	err = copier.Copy(&newCategoryBrands, categoryInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	for _, everyRecord := range newCategoryBrands {
		everyRecord.BrandId = req.BrandId
	}

	res = db.GetMysql().Create(&newCategoryBrands)
	if res.Error != nil {
		klog.Error(res.Error)
		if isConflict, kerr := utils.CheckResourceConflict(err, "Category or Brand conflicts."); isConflict {
			return nil, kerr
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = make([]*product.CategoryBrandInfo, 0)
	err = copier.Copy(&resp, newCategoryBrands)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
