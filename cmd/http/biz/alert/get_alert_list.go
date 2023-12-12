package alert

import (
	"context"

	alert "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetAlertListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAlertListService(Context context.Context, RequestContext *app.RequestContext) *GetAlertListService {
	return &GetAlertListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAlertListService) Do(req *alert.AlertFilter) (resp *alert.AlertListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
