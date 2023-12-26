package handler

import (
	"context"

	banner "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/banner"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddNewBannerService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewBannerService(Context context.Context, RequestContext *app.RequestContext) *AddNewBannerService {
	return &AddNewBannerService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewBannerService) Do(req *banner.NewBanner) (resp *banner.BannerInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
