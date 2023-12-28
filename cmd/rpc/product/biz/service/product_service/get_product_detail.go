package product_service

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type GetProductDetailService struct {
	ctx context.Context
}

// NewGetProductDetailService new GetProductDetailService
func NewGetProductDetailService(ctx context.Context) *GetProductDetailService {
	return &GetProductDetailService{ctx: ctx}
}

// Run create note info
func (s *GetProductDetailService) Run(req *common.Req) (resp *product.ProductDetailResp, err error) {
	// todo test joins
	productInfo := model.Product{}
	res := db.GetMysql().First(&productInfo, req.Id)
	if res.Error != nil {
		klog.Error(res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "The product is not existed.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.ProductDetailResp{}
	err = copier.Copy(resp, productInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	// Parsing the state field in product table from integer to  struct
	// And then assign the result value to the state field in resp
	parseStateToStruct(productInfo.State, resp)

	return resp, nil
}

// Convert the state of the product field
func parseStateToStruct(state model.ProductState, resp *product.ProductDetailResp) {
	resp.State = &product.ProductState{
		OnSale:         state&model.OnSale != 0,
		IsFreeShipping: state&model.IsFreeShipping != 0,
		IsNew:          state&model.IsNew != 0,
		IsHot:          state&model.IsHot != 0,
		IsRecommended:  state&model.IsRecommended != 0,
	}
}
