package brand_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetBrandDetailService struct {
	ctx context.Context
}

// NewGetBrandDetailService new GetBrandDetailService
func NewGetBrandDetailService(ctx context.Context) *GetBrandDetailService {
	return &GetBrandDetailService{ctx: ctx}
}

// Run create note info
func (s *GetBrandDetailService) Run(req *common.Req) (resp *product.BrandInfo, err error) {
	resp = &product.BrandInfo{}
	res := db.GetMysql().Table(model.Brand{}.TableName()).Where("id = ?", req.Id).Scan(resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("brand record is not found")
		return nil, kerrors.NewBizStatusError(code.NotFound, "The brand is not found.")
	}

	return resp, nil
}
