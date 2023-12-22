package user

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/mw"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	user "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCurUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCurUserInfoService(Context context.Context, RequestContext *app.RequestContext) *GetCurUserInfoService {
	return &GetCurUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCurUserInfoService) Do(req *common.Empty) (resp *user.UserInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	UserInfoInPayload, exists := h.RequestContext.Get(mw.IdentityKey)
	if !exists {
		hlog.Info("payload do not exist")
		return nil, err
	}
	hlog.Infof("user info in payload: %+v", UserInfoInPayload)

	userInfo, ok := UserInfoInPayload.(*model.User)
	if !ok {
		hlog.Info("failed to type assert")
		return nil, err
	}

	resp = &user.UserInfo{}
	err = copier.Copy(resp, userInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return resp, nil
}
