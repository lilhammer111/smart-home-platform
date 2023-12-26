package banner_service

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

type AddNewBannerService struct {
	ctx context.Context
}

// NewAddNewBannerService new AddNewBannerService
func NewAddNewBannerService(ctx context.Context) *AddNewBannerService {
	return &AddNewBannerService{ctx: ctx}
}

// Run create note info
func (s *AddNewBannerService) Run(req *product.NewBanner_) (resp *product.BannerInfo, err error) {
	bannerInfo := model.Banner{}
	err = copier.Copy(&bannerInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	err = db.GetMysql().Create(&bannerInfo).Error
	if err != nil {
		klog.Error(err)
		if isConflict, kerr := utils.CheckResourceConflict(err, "Banner index conflicts."); isConflict {
			return nil, kerr
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &product.BannerInfo{}
	err = copier.Copy(resp, bannerInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
