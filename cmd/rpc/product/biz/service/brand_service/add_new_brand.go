package brand_service

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

type AddNewBrandService struct {
	ctx context.Context
}

// NewAddNewBrandService new AddNewBrandService
func NewAddNewBrandService(ctx context.Context) *AddNewBrandService {
	return &AddNewBrandService{ctx: ctx}
}

// Run create note info
func (s *AddNewBrandService) Run(req *product.NewBrand_) (resp *product.BrandInfo, err error) {
	brandInfo := model.Brand{}
	err = copier.Copy(&brandInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	err = db.GetMysql().Create(&brandInfo).Error
	if err != nil {
		klog.Error(err)
		if isConflict, kerr := utils.CheckResourceConflict(err, "Brand name or logo conflicts."); isConflict {
			return nil, kerr
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
