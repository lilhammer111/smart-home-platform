package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/banner"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetAllBannersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAllBannersService(Context context.Context, RequestContext *app.RequestContext) *GetAllBannersService {
	return &GetAllBannersService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAllBannersService) Do(req *common.Empty) (resp *[]*banner.BannerInfo, err error) {
	bannerInfos, err := micro_product_cli.GetAllBanners(h.Context)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*banner.BannerInfo)
	err = copier.Copy(resp, bannerInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting all banners successes.")
	return resp, nil
}
