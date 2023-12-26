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

type AddNewBannerService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewBannerService(Context context.Context, RequestContext *app.RequestContext) *AddNewBannerService {
	return &AddNewBannerService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewBannerService) Do(req *banner.NewBanner) (resp *banner.BannerInfo, err error) {
	rpcBannerReq := product.NewBanner_{}
	err = copier.Copy(&rpcBannerReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	bannerInfo, err := micro_product_cli.AddNewBanner(h.Context, &rpcBannerReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &banner.BannerInfo{}
	err = copier.Copy(resp, bannerInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Adding a new brand successes.")
	return resp, nil
}
