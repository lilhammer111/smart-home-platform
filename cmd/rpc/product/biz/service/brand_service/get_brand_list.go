package brand_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetBrandListService struct {
	ctx context.Context
}

// NewGetBrandListService new GetBrandListService
func NewGetBrandListService(ctx context.Context) *GetBrandListService {
	return &GetBrandListService{ctx: ctx}
}

// Run create note info
func (s *GetBrandListService) Run(req *product.PageFilter) (resp []*product.BrandInfo, err error) {
	resp = make([]*product.BrandInfo, 0)
	res := db.GetMysql().Model(model.Brand{}.TableName()).Scopes(scope.Paginate(req.Page, req.Limit)).Find(resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.NotFound, "No brands.")
	}

	return resp, nil
}
