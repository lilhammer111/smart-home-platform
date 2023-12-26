package handler

import (
	"context"

	banner "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/banner"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetAllBannersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAllBannersService(Context context.Context, RequestContext *app.RequestContext) *GetAllBannersService {
	return &GetAllBannersService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAllBannersService) Do(req *common.Empty) (resp *[]*banner.BannerInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
