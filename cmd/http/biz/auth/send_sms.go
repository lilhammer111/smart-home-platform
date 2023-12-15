package auth

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/http/utils/user_micro_cli"
	"github.com/cloudwego/hertz/pkg/app"
)

type SendSmsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendSmsService(Context context.Context, RequestContext *app.RequestContext) *SendSmsService {
	return &SendSmsService{RequestContext: RequestContext, Context: Context}
}

func (h *SendSmsService) Do(req *auth.SendSmsReq) (resp *common_http.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	smsReq := micro_user.RpcSmsReq{Mobile: req.Mobile}
	_, err = user_micro_cli.DefaultClient().SendSMSViaAliyun(h.Context, &smsReq)
	return resp, err
}
