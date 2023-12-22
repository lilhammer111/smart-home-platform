package alert

import (
	"context"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteAlertService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteAlertService(Context context.Context, RequestContext *app.RequestContext) *DeleteAlertService {
	return &DeleteAlertService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteAlertService) Do(req *common.Req) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
