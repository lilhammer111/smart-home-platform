package banner_service

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

type GetAllBannersService struct {
	ctx context.Context
}

// NewGetAllBannersService new GetAllBannersService
func NewGetAllBannersService(ctx context.Context) *GetAllBannersService {
	return &GetAllBannersService{ctx: ctx}
}

// Run create note info
func (s *GetAllBannersService) Run() (resp []*product.BannerInfo, err error) {
	res := db.GetMysql().Model(&model.Banner{}).Find(&resp)
	if res.Error != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}
	if res.RowsAffected == 0 {
		klog.Info("no records")
		return nil, kerrors.NewBizStatusError(code.NotFound, "No banners")
	}

	return resp, nil
}
