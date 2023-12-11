package users

import (
	"context"

	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	users "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/users"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserListService(Context context.Context, RequestContext *app.RequestContext) *GetUserListService {
	return &GetUserListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserListService) Do(req *users.UsersFilter) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
