package handler

import (
	"context"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type DelteBannerService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDelteBannerService(Context context.Context, RequestContext *app.RequestContext) *DelteBannerService {
	return &DelteBannerService{RequestContext: RequestContext, Context: Context}
}

func (h *DelteBannerService) Do(req *common.Req) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
