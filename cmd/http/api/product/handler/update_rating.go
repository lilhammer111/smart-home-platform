package handler

import (
	"context"

	product "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateRatingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateRatingService(Context context.Context, RequestContext *app.RequestContext) *UpdateRatingService {
	return &UpdateRatingService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateRatingService) Do(req *product.RatingReq) (resp *product.RatingResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
