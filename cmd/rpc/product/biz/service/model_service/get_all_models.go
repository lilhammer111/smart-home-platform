package model_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetAllModelsService struct {
	ctx context.Context
}

// NewGetAllModelsService new GetAllModelsService
func NewGetAllModelsService(ctx context.Context) *GetAllModelsService {
	return &GetAllModelsService{ctx: ctx}
}

// Run create note info
func (s *GetAllModelsService) Run() (resp []*product.ModelInfo, err error) {
	res := db.GetMysql().Model(&model.Model{}).Find(&resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no records")
		return nil, kerrors.NewBizStatusError(code.NotFound, "No models")
	}

	return resp, nil
}
