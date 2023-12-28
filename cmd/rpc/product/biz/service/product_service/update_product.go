package product_service

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UpdateProductService struct {
	ctx context.Context
}

// NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.ProductInfo) (resp *product.ProductInfo, err error) {
	// todo bug of rating_id in product info
	err = db.GetMysql().First(&model.Product{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The product you'd like to update is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	productInfo := model.Product{}
	err = copier.Copy(&productInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	state := parseStateToInt(req.State.OnSale, req.State.IsNew, req.State.IsHot, req.State.IsFreeShipping, req.State.IsRecommended)
	productInfo.State = state

	res := db.GetMysql().Updates(&productInfo)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no updates")
		return nil, kerrors.NewBizStatusError(code.BadRequest, "No updates.")
	}

	resp = &product.ProductInfo{}
	err = copier.Copy(resp, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	return resp, nil
}
