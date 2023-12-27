package brand_service

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

type GetBrandDetailService struct {
	ctx context.Context
}

// NewGetBrandDetailService new GetBrandDetailService
func NewGetBrandDetailService(ctx context.Context) *GetBrandDetailService {
	return &GetBrandDetailService{ctx: ctx}
}

// Run create note info
func (s *GetBrandDetailService) Run(req *common.Req) (resp *product.BrandInfo, err error) {
	brandInfo := model.Brand{}
	res := db.GetMysql().Table(model.Brand{}.TableName()).Where("id = ?", req.Id).First(&brandInfo)
	if res.Error != nil {
		klog.Error(res.Error)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			klog.Infof("the category of id %d does not exist", req.Id)
			return nil, kerrors.NewBizStatusError(code.NotFound, "The brand is not found.")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.BrandInfo{}
	err = copier.Copy(resp, brandInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
