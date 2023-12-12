package feed_program

import (
	"context"

	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetSelfProgramListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetSelfProgramListService(Context context.Context, RequestContext *app.RequestContext) *GetSelfProgramListService {
	return &GetSelfProgramListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetSelfProgramListService) Do(req *int32) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
