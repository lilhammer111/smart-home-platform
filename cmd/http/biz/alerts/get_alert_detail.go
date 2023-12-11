package alerts

import (
	"context"

	req "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/req"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetAlertDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAlertDetailService(Context context.Context, RequestContext *app.RequestContext) *GetAlertDetailService {
	return &GetAlertDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAlertDetailService) Do(req *req.IdReq) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
