package auth

import (
	"context"

	auth "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type PwdLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPwdLoginService(Context context.Context, RequestContext *app.RequestContext) *PwdLoginService {
	return &PwdLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *PwdLoginService) Do(req *auth.PwdLoginReq) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
