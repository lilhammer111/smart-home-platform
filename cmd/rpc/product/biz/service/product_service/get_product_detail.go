package product_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetProductDetailService struct {
	ctx context.Context
}

// NewGetProductDetailService new GetProductDetailService
func NewGetProductDetailService(ctx context.Context) *GetProductDetailService {
	return &GetProductDetailService{ctx: ctx}
}

// Run create note info
func (s *GetProductDetailService) Run(req *common.Req) (resp *product.ProductDetail, err error) {
	resp = &product.ProductDetail{}
	// todo test joins
	res := db.GetMysql().Table(model.Product{}.TableName()).
		Where("id = ?", req.Id).
		Joins(model.Brand{}.TableName()).
		Joins(model.Category{}.TableName()).
		Joins(model.Model{}.TableName()).
		Find(resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info(utils.InfoNotFound)
		return nil, kerrors.NewBizStatusError(code.NotFound, "The product is not found.")
	}

	return resp, nil
}
