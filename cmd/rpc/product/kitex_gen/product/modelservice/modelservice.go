// Code generated by Kitex v0.8.0. DO NOT EDIT.

package modelservice

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return modelServiceServiceInfo
}

var modelServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ModelService"
	handlerType := (*product.ModelService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetAllModels":   kitex.NewMethodInfo(getAllModelsHandler, newModelServiceGetAllModelsArgs, newModelServiceGetAllModelsResult, false),
		"GetModelDetail": kitex.NewMethodInfo(getModelDetailHandler, newModelServiceGetModelDetailArgs, newModelServiceGetModelDetailResult, false),
		"AddNewModel":    kitex.NewMethodInfo(addNewModelHandler, newModelServiceAddNewModelArgs, newModelServiceAddNewModelResult, false),
		"DeleteModel":    kitex.NewMethodInfo(deleteModelHandler, newModelServiceDeleteModelArgs, newModelServiceDeleteModelResult, false),
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

func getAllModelsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*product.ModelServiceGetAllModelsResult)
	success, err := handler.(product.ModelService).GetAllModels(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newModelServiceGetAllModelsArgs() interface{} {
	return product.NewModelServiceGetAllModelsArgs()
}

func newModelServiceGetAllModelsResult() interface{} {
	return product.NewModelServiceGetAllModelsResult()
}

func getModelDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ModelServiceGetModelDetailArgs)
	realResult := result.(*product.ModelServiceGetModelDetailResult)
	success, err := handler.(product.ModelService).GetModelDetail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newModelServiceGetModelDetailArgs() interface{} {
	return product.NewModelServiceGetModelDetailArgs()
}

func newModelServiceGetModelDetailResult() interface{} {
	return product.NewModelServiceGetModelDetailResult()
}

func addNewModelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ModelServiceAddNewModelArgs)
	realResult := result.(*product.ModelServiceAddNewModelResult)
	success, err := handler.(product.ModelService).AddNewModel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newModelServiceAddNewModelArgs() interface{} {
	return product.NewModelServiceAddNewModelArgs()
}

func newModelServiceAddNewModelResult() interface{} {
	return product.NewModelServiceAddNewModelResult()
}

func deleteModelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ModelServiceDeleteModelArgs)
	realResult := result.(*product.ModelServiceDeleteModelResult)
	success, err := handler.(product.ModelService).DeleteModel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newModelServiceDeleteModelArgs() interface{} {
	return product.NewModelServiceDeleteModelArgs()
}

func newModelServiceDeleteModelResult() interface{} {
	return product.NewModelServiceDeleteModelResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetAllModels(ctx context.Context) (r []*product.ModelInfo, err error) {
	var _args product.ModelServiceGetAllModelsArgs
	var _result product.ModelServiceGetAllModelsResult
	if err = p.c.Call(ctx, "GetAllModels", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetModelDetail(ctx context.Context, req *common.Req) (r *product.ModelInfo, err error) {
	var _args product.ModelServiceGetModelDetailArgs
	_args.Req = req
	var _result product.ModelServiceGetModelDetailResult
	if err = p.c.Call(ctx, "GetModelDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddNewModel(ctx context.Context, req *product.NewModel_) (r *product.ModelInfo, err error) {
	var _args product.ModelServiceAddNewModelArgs
	_args.Req = req
	var _result product.ModelServiceAddNewModelResult
	if err = p.c.Call(ctx, "AddNewModel", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteModel(ctx context.Context, req *common.Req) (r *common.Empty, err error) {
	var _args product.ModelServiceDeleteModelArgs
	_args.Req = req
	var _result product.ModelServiceDeleteModelResult
	if err = p.c.Call(ctx, "DeleteModel", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
