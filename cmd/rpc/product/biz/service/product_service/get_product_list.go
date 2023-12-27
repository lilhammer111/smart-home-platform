package product_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetProductListService struct {
	ctx context.Context
}

// NewGetProductListService new GetProductListService
func NewGetProductListService(ctx context.Context) *GetProductListService {
	return &GetProductListService{ctx: ctx}
}

// Run create note info
func (s *GetProductListService) Run(req *product.ProductFilter) (resp []*product.ProductDetail, err error) {
	tx := db.GetMysql()
	if req.State != nil {
		tx = tx.Where("state = ?", req.State)
	}
	if req.BrandId != nil {
		tx = tx.Where("brand_id = ?", req.BrandId)
	}
	if req.CategoryId != nil {
		tx = tx.Where("category_id = ?", req.CategoryId)
	}
	if req.Sorts != nil {
		for _, sortCond := range req.Sorts {
			tx = tx.Order(sortCond)
		}
	}

	resp = make([]*product.ProductDetail, 0)
	res := tx.Scopes(scope.Paginate(req.Page, req.Limit)).
		Joins(model.Brand{}.TableName()).
		Joins(model.Category{}.TableName()).
		Joins(model.Model{}.TableName()).
		Find(&resp)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info(utils.InfoNotFound)
		return nil, kerrors.NewBizStatusError(code.NotFound, "No eligible products.")
	}

	return resp, nil
}
