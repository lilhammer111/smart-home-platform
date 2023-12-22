package alert

import (
	"context"

	alert "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetAlertDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAlertDetailService(Context context.Context, RequestContext *app.RequestContext) *GetAlertDetailService {
	return &GetAlertDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAlertDetailService) Do(req *common.Req) (resp *alert.AlertInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
