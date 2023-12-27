package category_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetCategoryListService struct {
	ctx context.Context
}

// NewGetCategoryListService new GetCategoryListService
func NewGetCategoryListService(ctx context.Context) *GetCategoryListService {
	return &GetCategoryListService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryListService) Run(req *product.PageFilter) (resp []*product.CategoryInfo, err error) {
	resp = make([]*product.CategoryInfo, 0)
	res := db.GetMysql().Model(model.Brand{}.TableName()).Scopes(scope.Paginate(req.Page, req.Limit)).Find(resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.NotFound, "No categories.")
	}

	return resp, nil
}
