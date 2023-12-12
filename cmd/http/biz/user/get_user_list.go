package user

import (
	"context"

	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	user "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserListService(Context context.Context, RequestContext *app.RequestContext) *GetUserListService {
	return &GetUserListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserListService) Do(req *user.UsersFilter) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
