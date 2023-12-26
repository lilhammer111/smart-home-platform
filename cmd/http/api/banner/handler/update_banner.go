package handler

import (
	"context"

	banner "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/banner"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateBannerService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateBannerService(Context context.Context, RequestContext *app.RequestContext) *UpdateBannerService {
	return &UpdateBannerService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateBannerService) Do(req *banner.BannerInfo) (resp *banner.BannerInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
