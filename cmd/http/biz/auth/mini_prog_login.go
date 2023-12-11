package auth

import (
	"context"

	auth "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/auth"
	resp "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type MiniProgLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMiniProgLoginService(Context context.Context, RequestContext *app.RequestContext) *MiniProgLoginService {
	return &MiniProgLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *MiniProgLoginService) Do(req *auth.MobileRegisterReq) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
