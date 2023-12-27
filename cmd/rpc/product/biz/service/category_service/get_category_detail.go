package category_service

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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
	categoryInfo := model.Category{}
	res := db.GetMysql().Table(model.Category{}.TableName()).Where("id = ?", req.Id).First(&categoryInfo)
	if res.Error != nil {
		klog.Error(res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			klog.Infof("the category of id %d does not exist", req.Id)
			return nil, kerrors.NewBizStatusError(code.NotFound, "The category is not found.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.CategoryInfo{}
	err = copier.Copy(resp, categoryInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
