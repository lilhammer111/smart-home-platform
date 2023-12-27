package product_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type AddNewProductService struct {
	ctx context.Context
}

// NewAddNewProductService new AddNewProductService
func NewAddNewProductService(ctx context.Context) *AddNewProductService {
	return &AddNewProductService{ctx: ctx}
}

// Run create note info
func (s *AddNewProductService) Run(req *product.NewProduct_) (resp *product.ProductInfo, err error) {
	productInfo := model.Product{}
	err = copier.Copy(&productInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	err = db.GetMysql().Create(&productInfo).Error
	if err != nil {
		klog.Error(err)
		// todo handle the different situation of product name conflict and rating product_id conflict
		if isConflict, kerr := utils.CheckResourceConflict(err, "Product name conflicts."); isConflict {
			return nil, kerr
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.ProductInfo{}
	err = copier.Copy(resp, productInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
