package service

import (
	"context"
	"fmt"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"

	"github.com/tidwall/gjson"
)

type RequestOpenIdService struct {
	ctx context.Context
} // NewRequestOpenIdService new RequestOpenIdService
func NewRequestOpenIdService(ctx context.Context) *RequestOpenIdService {
	return &RequestOpenIdService{ctx: ctx}
}

// Run create note info
func (s *RequestOpenIdService) Run(req *micro_user.RpcRequestOpenIdReq) (resp *micro_user.RpcRequestOpenIdResp, err error) {
	// Finish your business logic.
	// invoke wx mini program api
	cli := &client.Client{}
	cli, err = client.NewClient()
	if err != nil {
		klog.Errorf("failed to create a http client: %s", err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	httpReq, httpResp := protocol.AcquireRequest(), protocol.AcquireResponse()
	httpReq.SetRequestURI("https://api.weixin.qq.com/sns/jscode2session")
	httpReq.SetMethod(http.MethodGet)
	httpReq.SetQueryString(
		fmt.Sprintf("appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
			req.Appid,
			req.Secret,
			req.JsCode,
		),
	)
	err = cli.Do(context.Background(), httpReq, httpResp)
	if err != nil {
		klog.Errorf("failed to do request: %s", err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}
	//type MiniProgResp struct {
	//	Openid     string `json:"openid,omitempty"`
	//	SessionKey string `json:"session_key,omitempty"`
	//	UnionId    string `json:"unionid,omitempty"`
	//	ErrCode    int32  `json:"errcode,omitempty"`
	//	ErrMsg     string `json:"errmsg,omitempty"`
	//}
	//miniProgResp := &MiniProgResp{}
	//err = json.Unmarshal(httpResp.Body(), &miniProgResp)
	//if err != nil {
	//	hlog.Errorf("failed to unmarshal miniProgResp: %s", err)
	//	return nil, ErrMiniProgLoginFailed
	//}
	openid := gjson.Get(string(httpResp.Body()), "openid").String()
	return &micro_user.RpcRequestOpenIdResp{OpenId: openid}, nil
}
