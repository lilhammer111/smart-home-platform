// Code generated by Kitex v0.8.0. DO NOT EDIT.

package brand

import (
	"context"
	brand "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/brand"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return brandServiceInfo
}

var brandServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Brand"
	handlerType := (*brand.Brand)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetBrandList":                 kitex.NewMethodInfo(getBrandListHandler, newBrandGetBrandListArgs, newBrandGetBrandListResult, false),
		"GetRelatedBrandsByCategoryId": kitex.NewMethodInfo(getRelatedBrandsByCategoryIdHandler, newBrandGetRelatedBrandsByCategoryIdArgs, newBrandGetRelatedBrandsByCategoryIdResult, false),
		"GetBrandDetail":               kitex.NewMethodInfo(getBrandDetailHandler, newBrandGetBrandDetailArgs, newBrandGetBrandDetailResult, false),
		"AddNewBrand":                  kitex.NewMethodInfo(addNewBrandHandler, newBrandAddNewBrandArgs, newBrandAddNewBrandResult, false),
		"UpdateBrand":                  kitex.NewMethodInfo(updateBrandHandler, newBrandUpdateBrandArgs, newBrandUpdateBrandResult, false),
		"DeleteBrand":                  kitex.NewMethodInfo(deleteBrandHandler, newBrandDeleteBrandArgs, newBrandDeleteBrandResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "brand",
		"ServiceFilePath": `../../../static/idl/http/brand.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func getBrandListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*brand.BrandGetBrandListArgs)
	realResult := result.(*brand.BrandGetBrandListResult)
	success, err := handler.(brand.Brand).GetBrandList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBrandGetBrandListArgs() interface{} {
	return brand.NewBrandGetBrandListArgs()
}

func newBrandGetBrandListResult() interface{} {
	return brand.NewBrandGetBrandListResult()
}

func getRelatedBrandsByCategoryIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*brand.BrandGetRelatedBrandsByCategoryIdArgs)
	realResult := result.(*brand.BrandGetRelatedBrandsByCategoryIdResult)
	success, err := handler.(brand.Brand).GetRelatedBrandsByCategoryId(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBrandGetRelatedBrandsByCategoryIdArgs() interface{} {
	return brand.NewBrandGetRelatedBrandsByCategoryIdArgs()
}

func newBrandGetRelatedBrandsByCategoryIdResult() interface{} {
	return brand.NewBrandGetRelatedBrandsByCategoryIdResult()
}

func getBrandDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*brand.BrandGetBrandDetailArgs)
	realResult := result.(*brand.BrandGetBrandDetailResult)
	success, err := handler.(brand.Brand).GetBrandDetail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBrandGetBrandDetailArgs() interface{} {
	return brand.NewBrandGetBrandDetailArgs()
}

func newBrandGetBrandDetailResult() interface{} {
	return brand.NewBrandGetBrandDetailResult()
}

func addNewBrandHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*brand.BrandAddNewBrandArgs)
	realResult := result.(*brand.BrandAddNewBrandResult)
	success, err := handler.(brand.Brand).AddNewBrand(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBrandAddNewBrandArgs() interface{} {
	return brand.NewBrandAddNewBrandArgs()
}

func newBrandAddNewBrandResult() interface{} {
	return brand.NewBrandAddNewBrandResult()
}

func updateBrandHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*brand.BrandUpdateBrandArgs)
	realResult := result.(*brand.BrandUpdateBrandResult)
	success, err := handler.(brand.Brand).UpdateBrand(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBrandUpdateBrandArgs() interface{} {
	return brand.NewBrandUpdateBrandArgs()
}

func newBrandUpdateBrandResult() interface{} {
	return brand.NewBrandUpdateBrandResult()
}

func deleteBrandHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*brand.BrandDeleteBrandArgs)
	realResult := result.(*brand.BrandDeleteBrandResult)
	success, err := handler.(brand.Brand).DeleteBrand(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBrandDeleteBrandArgs() interface{} {
	return brand.NewBrandDeleteBrandArgs()
}

func newBrandDeleteBrandResult() interface{} {
	return brand.NewBrandDeleteBrandResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetBrandList(ctx context.Context, req *common.PageFilter) (r []*brand.BrandListResp, err error) {
	var _args brand.BrandGetBrandListArgs
	_args.Req = req
	var _result brand.BrandGetBrandListResult
	if err = p.c.Call(ctx, "GetBrandList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetRelatedBrandsByCategoryId(ctx context.Context, req *brand.BrandByCatReq) (r []*brand.BrandInfo, err error) {
	var _args brand.BrandGetRelatedBrandsByCategoryIdArgs
	_args.Req = req
	var _result brand.BrandGetRelatedBrandsByCategoryIdResult
	if err = p.c.Call(ctx, "GetRelatedBrandsByCategoryId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetBrandDetail(ctx context.Context, req *common.Req) (r *brand.BrandInfo, err error) {
	var _args brand.BrandGetBrandDetailArgs
	_args.Req = req
	var _result brand.BrandGetBrandDetailResult
	if err = p.c.Call(ctx, "GetBrandDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddNewBrand(ctx context.Context, req *brand.NewBrand_) (r *brand.BrandInfo, err error) {
	var _args brand.BrandAddNewBrandArgs
	_args.Req = req
	var _result brand.BrandAddNewBrandResult
	if err = p.c.Call(ctx, "AddNewBrand", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBrand(ctx context.Context, req *brand.BrandInfo) (r *brand.BrandInfo, err error) {
	var _args brand.BrandUpdateBrandArgs
	_args.Req = req
	var _result brand.BrandUpdateBrandResult
	if err = p.c.Call(ctx, "UpdateBrand", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteBrand(ctx context.Context, req *common.Req) (r *common.Empty, err error) {
	var _args brand.BrandDeleteBrandArgs
	_args.Req = req
	var _result brand.BrandDeleteBrandResult
	if err = p.c.Call(ctx, "DeleteBrand", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
