package feed_program

import (
	"context"

	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteProgramService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProgramService(Context context.Context, RequestContext *app.RequestContext) *DeleteProgramService {
	return &DeleteProgramService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProgramService) Do(req *int32) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
