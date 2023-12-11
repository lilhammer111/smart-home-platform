package users

import (
	"context"

	resp "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/resp"
	users "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/users"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Do(req *users.UserInfo) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
