package user

import (
	"context"
	rpcUser "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_user_cli"
	"github.com/jinzhu/copier"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserListService(Context context.Context, RequestContext *app.RequestContext) *GetUserListService {
	return &GetUserListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserListService) Do(req *user.UsersFilter) (resp *[]*user.UserInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcReq := rpcUser.UsersFilter{}
	if err = copier.Copy(&rpcReq, req); err != nil {
		return nil, err
	}

	userInfoList, err := micro_user_cli.QueryUsersWithFilter(h.Context, &rpcReq)
	if err != nil {
		return nil, err
	}

	respList := make([]*user.UserInfo, 0)
	for _, userInfo := range userInfoList {
		respInfo := &user.UserInfo{}
		if err = copier.Copy(respInfo, &userInfo); err != nil {
			return nil, err
		}
		respList = append(respList, respInfo)
	}
	return &respList, nil
}
