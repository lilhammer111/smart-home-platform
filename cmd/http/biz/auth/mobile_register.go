package auth

import (
	"context"
	auth "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_user_cli"
	"github.com/cloudwego/hertz/pkg/app"
)

type MobileRegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMobileRegisterService(Context context.Context, RequestContext *app.RequestContext) *MobileRegisterService {
	return &MobileRegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *MobileRegisterService) Do(req *auth.MobileRegisterReq) (resp *auth.AuthInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// check sms code
	isCorrect, err := micro_user_cli.DefaultClient().VerifySmsCode(h.Context, req.Mobile, req.SmsCode)
	if err != nil {
		return nil, err
	}
	// create user
	userInfo, err := micro_user_cli.CreateUser(h.Context, req)
	resp = auth.AuthInfoResp{}
	return
}
