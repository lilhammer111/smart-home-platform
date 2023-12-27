package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"

	"github.com/cloudwego/hertz/pkg/app"
)

type MiniProgLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMiniProgLoginService(Context context.Context, RequestContext *app.RequestContext) *MiniProgLoginService {
	return &MiniProgLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *MiniProgLoginService) Do(req *auth.MiniProgLoginReq) (resp *auth.AuthInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}