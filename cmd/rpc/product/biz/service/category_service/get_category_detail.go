package category_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetCategoryDetailService struct {
	ctx context.Context
}

// NewGetCategoryDetailService new GetCategoryDetailService
func NewGetCategoryDetailService(ctx context.Context) *GetCategoryDetailService {
	return &GetCategoryDetailService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryDetailService) Run(req *common.Req) (resp *product.CategoryInfo, err error) {
	resp = &product.CategoryInfo{}
	res := db.GetMysql().Table(model.Brand{}.TableName()).Where("id = ?", req.Id).Scan(resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info(utils.InfoNotFound)
		return nil, kerrors.NewBizStatusError(code.NotFound, "The category is not found.")
	}

	return resp, nil
}
