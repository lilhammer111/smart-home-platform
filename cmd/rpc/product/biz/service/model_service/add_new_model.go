package model_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/utils"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type AddNewModelService struct {
	ctx context.Context
}

// NewAddNewModelService new AddNewModelService
func NewAddNewModelService(ctx context.Context) *AddNewModelService {
	return &AddNewModelService{ctx: ctx}
}

// Run create note info
func (s *AddNewModelService) Run(req *product.NewModel_) (resp *product.ModelInfo, err error) {
	modelInfo := model.Model{}
	err = copier.Copy(&modelInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	err = db.GetMysql().Create(&modelInfo).Error
	if err != nil {
		klog.Error(err)
		if isConflict, kerr := utils.CheckResourceConflict(err, "Banner name conflicts."); isConflict {
			return nil, kerr
		}

		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.ModelInfo{}
	err = copier.Copy(resp, modelInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
