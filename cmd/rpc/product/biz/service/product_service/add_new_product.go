package product_service

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

type AddNewProductService struct {
	ctx context.Context
}

// NewAddNewProductService new AddNewProductService
func NewAddNewProductService(ctx context.Context) *AddNewProductService {
	return &AddNewProductService{ctx: ctx}
}

// Run create note info
func (s *AddNewProductService) Run(req *product.AddProductReq) (resp *product.ProductInfo, err error) {
	// Check if the brand id exists
	err = db.GetMysql().Select("id").First(&model.Brand{}, req.BrandId).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			klog.Infof("the brand id %d is not existed", req.BrandId)
			return nil, kerrors.NewBizStatusError(code.BadRequest, "Adding the product failed. Because the brand id is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	// Check if the category id exists
	err = db.GetMysql().Select("id").First(&model.Category{}, req.CategoryId).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			klog.Info("the category is not existed")
			return nil, kerrors.NewBizStatusError(code.BadRequest, "Adding the product failed. Because the category is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	productInfo := model.Product{}
	// todo test
	err = copier.Copy(&productInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	state := parseStateToInt(&req.State.OnSale, &req.State.IsNew, &req.State.IsHot, &req.State.IsFreeShipping, &req.State.IsRecommended)
	productInfo.State = state

	err = db.GetMysql().Create(&productInfo).Error
	if err != nil {
		klog.Error(err)
		if isConflict, kerr := utils.CheckResourceConflict(err, "Product name conflicts."); isConflict {
			return nil, kerr
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.ProductInfo{}
	err = copier.Copy(resp, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	resp.Id = productInfo.ID

	return resp, nil
}
