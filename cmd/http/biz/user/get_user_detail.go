package user

import (
	"context"
	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	user "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	std "git.zqbjj.top/pet/services/cmd/http/kitex_gen/standard"
	"git.zqbjj.top/pet/services/cmd/http/utils/rpc_client/user_micro"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetUserDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserDetailService(Context context.Context, RequestContext *app.RequestContext) *GetUserDetailService {
	return &GetUserDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserDetailService) Do(req *standard.Req) (resp *user.UserInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcReq := &std.Req{Id: req.Id}

	rpcResp, err := user_micro.DefaultClient().FindUser(h.Context, rpcReq)
	if err != nil {
		hlog.Errorf("rpc: FindUser: failed to get user detail: %s", err)
		return nil, err
	}

	err = copier.CopyWithOption(resp.Data, rpcResp, copier.Option{DeepCopy: true})
	if err != nil {
		hlog.Errorf("http: GetUserDetail: failed to copy fields: %s", err)
		return nil, err
	}
	resp = &user.UserInfoResp{
		Meta: &standard.Resp{
			Success: true,
			Code:    200,
			Message: "OK",
		},
		//Data: &user.UserInfo{
		//	Id:       rpcResp.Id,
		//	Age:      rpcResp.Age,
		//	Gender:   rpcResp.Gender,
		//	Mobile:   rpcResp.Mobile,
		//	Profile:  rpcResp.Profile,
		//	Username: rpcResp.Username,
		//	Email:    rpcResp.Email,
		//	Avatar:   rpcResp.Avatar,
		//},
	}
	return
}
