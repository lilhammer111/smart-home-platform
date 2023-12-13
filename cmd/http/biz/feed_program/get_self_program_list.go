package feed_program

import (
	"context"

	common_http "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common_http"
	feed_program "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/feed_program"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetSelfProgramListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetSelfProgramListService(Context context.Context, RequestContext *app.RequestContext) *GetSelfProgramListService {
	return &GetSelfProgramListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetSelfProgramListService) Do(req *common_http.Req) (resp *feed_program.ProgramListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
