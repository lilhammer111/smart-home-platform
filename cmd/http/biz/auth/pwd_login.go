package auth

import (
	"context"

	auth "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type PwdLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPwdLoginService(Context context.Context, RequestContext *app.RequestContext) *PwdLoginService {
	return &PwdLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *PwdLoginService) Do(req *auth.PwdLoginReq) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
