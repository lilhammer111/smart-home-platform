package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetBrandListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetBrandListService(Context context.Context, RequestContext *app.RequestContext) *GetBrandListService {
	return &GetBrandListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetBrandListService) Do(req *common.PageFilter) (resp *[]*brand.BrandInfo, err error) {
	pageFilter := rpcCommon.PageFilter{}
	err = copier.Copy(&pageFilter, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	brandInfos, err := micro_product_cli.GetBrandList(h.Context, &pageFilter)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*brand.BrandInfo)
	err = copier.Copy(resp, brandInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting brands list successes.")
	return resp, nil
}
