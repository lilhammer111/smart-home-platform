package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/banner"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type UpdateBannerService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateBannerService(Context context.Context, RequestContext *app.RequestContext) *UpdateBannerService {
	return &UpdateBannerService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateBannerService) Do(req *banner.BannerInfo) (resp *banner.BannerInfo, err error) {
	bannerReq := product.BannerInfo{}
	err = copier.Copy(&bannerReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	updatedBannerInfos, err := micro_product_cli.UpdateBanner(h.Context, &bannerReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &banner.BannerInfo{}
	err = copier.Copy(resp, updatedBannerInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Updating banner info successes.")
	return resp, nil
}
