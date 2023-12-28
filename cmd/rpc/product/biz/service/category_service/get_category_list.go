package category_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/lilhammer111/micro-kit/model/scope"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type GetCategoryListService struct {
	ctx context.Context
}

// NewGetCategoryListService new GetCategoryListService
func NewGetCategoryListService(ctx context.Context) *GetCategoryListService {
	return &GetCategoryListService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryListService) Run(req *common.PageFilter) (resp []*product.CategoryInfo, err error) {
	categoryInfos := make([]model.Category, 0)
	res := db.GetMysql().Model(&model.Category{}).Scopes(scope.Paginate(req.Page, req.Limit)).Find(&categoryInfos)
	if res.Error != nil {
		klog.Error(res.Error)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("the record is not found")
		return nil, kerrors.NewBizStatusError(code.NotFound, "No categories.")
	}
	resp = make([]*product.CategoryInfo, 0)
	err = copier.Copy(&resp, categoryInfos)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
