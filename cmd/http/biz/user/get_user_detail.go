package user

import (
	"context"

	common_http "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	user "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserDetailService(Context context.Context, RequestContext *app.RequestContext) *GetUserDetailService {
	return &GetUserDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserDetailService) Do(req *common_http.Req) (resp *user.UserInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
