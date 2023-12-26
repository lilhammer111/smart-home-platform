package model_service

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

type GetModelDetailService struct {
	ctx context.Context
}

// NewGetModelDetailService new GetModelDetailService
func NewGetModelDetailService(ctx context.Context) *GetModelDetailService {
	return &GetModelDetailService{ctx: ctx}
}

// Run create note info
func (s *GetModelDetailService) Run(req *common.Req) (resp *product.ModelInfo, err error) {
	resp = &product.ModelInfo{}
	res := db.GetMysql().Table(model.Model{}.TableName()).Where("id = ?", req.Id).Scan(resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("model record is not found")
		return nil, kerrors.NewBizStatusError(code.NotFound, "The model is not found.")
	}

	return resp, nil
}
