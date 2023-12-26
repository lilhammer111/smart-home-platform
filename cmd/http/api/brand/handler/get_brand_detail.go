package handler

import (
	"context"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetBrandDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetBrandDetailService(Context context.Context, RequestContext *app.RequestContext) *GetBrandDetailService {
	return &GetBrandDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetBrandDetailService) Do(req *common.Req) (resp *brand.BrandInfo, err error) {
	brandInfo, err := micro_product_cli.GetBrandDetail(h.Context, &rpcCommon.Req{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &brand.BrandInfo{}
	err = copier.Copy(resp, brandInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting brand details successes.")
	return resp, nil
}
