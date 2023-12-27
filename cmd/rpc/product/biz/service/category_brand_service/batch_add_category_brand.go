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

	categoryInfos := make([]model.Category, 0)
	// todo test_0001
	res := db.GetMysql().Find(&categoryInfos, req.CategoryId)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no categories.")
		return nil, kerrors.NewBizStatusError(code.NotFound, "None of categories exist.")
	}

	newCategoryBrandInfos := make([]model.CategoryBrand, 0)
	for _, category := range categoryInfos {
		newCategoryBrandInfos = append(newCategoryBrandInfos, model.CategoryBrand{
			BrandId:    req.BrandId,
			CategoryId: category.Id,
		})
	}

	res = db.GetMysql().Create(&newCategoryBrandInfos)
	if res.Error != nil {
		klog.Error(res.Error)
		if isConflict, kerr := utils.CheckResourceConflict(res.Error, "Category or Brand conflicts."); isConflict {
			return nil, kerr
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = make([]*product.CategoryBrandInfo, 0)
	err = copier.Copy(&resp, newCategoryBrandInfos)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
