// Code generated by Kitex v0.8.0. DO NOT EDIT.

package bannerservice

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return bannerServiceServiceInfo
}

var bannerServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "BannerService"
	handlerType := (*product.BannerService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetAllBanners": kitex.NewMethodInfo(getAllBannersHandler, newBannerServiceGetAllBannersArgs, newBannerServiceGetAllBannersResult, false),
		"AddNewBanner":  kitex.NewMethodInfo(addNewBannerHandler, newBannerServiceAddNewBannerArgs, newBannerServiceAddNewBannerResult, false),
		"UpdateBanner":  kitex.NewMethodInfo(updateBannerHandler, newBannerServiceUpdateBannerArgs, newBannerServiceUpdateBannerResult, false),
		"DeleteBanner":  kitex.NewMethodInfo(deleteBannerHandler, newBannerServiceDeleteBannerArgs, newBannerServiceDeleteBannerResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "product",
		"ServiceFilePath": `../../../static/idl/rpc/product.thrift`,
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

func getAllBannersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*product.BannerServiceGetAllBannersResult)
	success, err := handler.(product.BannerService).GetAllBanners(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBannerServiceGetAllBannersArgs() interface{} {
	return product.NewBannerServiceGetAllBannersArgs()
}

func newBannerServiceGetAllBannersResult() interface{} {
	return product.NewBannerServiceGetAllBannersResult()
}

func addNewBannerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.BannerServiceAddNewBannerArgs)
	realResult := result.(*product.BannerServiceAddNewBannerResult)
	success, err := handler.(product.BannerService).AddNewBanner(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBannerServiceAddNewBannerArgs() interface{} {
	return product.NewBannerServiceAddNewBannerArgs()
}

func newBannerServiceAddNewBannerResult() interface{} {
	return product.NewBannerServiceAddNewBannerResult()
}

func updateBannerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.BannerServiceUpdateBannerArgs)
	realResult := result.(*product.BannerServiceUpdateBannerResult)
	success, err := handler.(product.BannerService).UpdateBanner(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBannerServiceUpdateBannerArgs() interface{} {
	return product.NewBannerServiceUpdateBannerArgs()
}

func newBannerServiceUpdateBannerResult() interface{} {
	return product.NewBannerServiceUpdateBannerResult()
}

func deleteBannerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.BannerServiceDeleteBannerArgs)
	realResult := result.(*product.BannerServiceDeleteBannerResult)
	success, err := handler.(product.BannerService).DeleteBanner(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBannerServiceDeleteBannerArgs() interface{} {
	return product.NewBannerServiceDeleteBannerArgs()
}

func newBannerServiceDeleteBannerResult() interface{} {
	return product.NewBannerServiceDeleteBannerResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetAllBanners(ctx context.Context) (r []*product.BannerInfo, err error) {
	var _args product.BannerServiceGetAllBannersArgs
	var _result product.BannerServiceGetAllBannersResult
	if err = p.c.Call(ctx, "GetAllBanners", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddNewBanner(ctx context.Context, req *product.NewBanner_) (r *product.BannerInfo, err error) {
	var _args product.BannerServiceAddNewBannerArgs
	_args.Req = req
	var _result product.BannerServiceAddNewBannerResult
	if err = p.c.Call(ctx, "AddNewBanner", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBanner(ctx context.Context, req *product.BannerInfo) (r *product.BannerInfo, err error) {
	var _args product.BannerServiceUpdateBannerArgs
	_args.Req = req
	var _result product.BannerServiceUpdateBannerResult
	if err = p.c.Call(ctx, "UpdateBanner", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteBanner(ctx context.Context, req *common.Req) (r *common.Empty, err error) {
	var _args product.BannerServiceDeleteBannerArgs
	_args.Req = req
	var _result product.BannerServiceDeleteBannerResult
	if err = p.c.Call(ctx, "DeleteBanner", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
