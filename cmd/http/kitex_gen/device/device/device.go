// Code generated by Kitex v0.6.1. DO NOT EDIT.

package device

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	device "git.zqbjj.top/pet/services/cmd/http/kitex_gen/device"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return deviceServiceInfo
}

var deviceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "device"
	handlerType := (*device.Device)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetDeviceList":    kitex.NewMethodInfo(getDeviceListHandler, newDeviceGetDeviceListArgs, newDeviceGetDeviceListResult, false),
		"GetDeviceDetail":  kitex.NewMethodInfo(getDeviceDetailHandler, newDeviceGetDeviceDetailArgs, newDeviceGetDeviceDetailResult, false),
		"UpdateDeviceInfo": kitex.NewMethodInfo(updateDeviceInfoHandler, newDeviceUpdateDeviceInfoArgs, newDeviceUpdateDeviceInfoResult, false),
		"BindDevice":       kitex.NewMethodInfo(bindDeviceHandler, newDeviceBindDeviceArgs, newDeviceBindDeviceResult, false),
		"UnbindDevice":     kitex.NewMethodInfo(unbindDeviceHandler, newDeviceUnbindDeviceArgs, newDeviceUnbindDeviceResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "device",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func getDeviceListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*device.DeviceGetDeviceListArgs)
	realResult := result.(*device.DeviceGetDeviceListResult)
	success, err := handler.(device.Device).GetDeviceList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDeviceGetDeviceListArgs() interface{} {
	return device.NewDeviceGetDeviceListArgs()
}

func newDeviceGetDeviceListResult() interface{} {
	return device.NewDeviceGetDeviceListResult()
}

func getDeviceDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*device.DeviceGetDeviceDetailArgs)
	realResult := result.(*device.DeviceGetDeviceDetailResult)
	success, err := handler.(device.Device).GetDeviceDetail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDeviceGetDeviceDetailArgs() interface{} {
	return device.NewDeviceGetDeviceDetailArgs()
}

func newDeviceGetDeviceDetailResult() interface{} {
	return device.NewDeviceGetDeviceDetailResult()
}

func updateDeviceInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*device.DeviceUpdateDeviceInfoArgs)
	realResult := result.(*device.DeviceUpdateDeviceInfoResult)
	success, err := handler.(device.Device).UpdateDeviceInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDeviceUpdateDeviceInfoArgs() interface{} {
	return device.NewDeviceUpdateDeviceInfoArgs()
}

func newDeviceUpdateDeviceInfoResult() interface{} {
	return device.NewDeviceUpdateDeviceInfoResult()
}

func bindDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*device.DeviceBindDeviceArgs)
	realResult := result.(*device.DeviceBindDeviceResult)
	success, err := handler.(device.Device).BindDevice(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDeviceBindDeviceArgs() interface{} {
	return device.NewDeviceBindDeviceArgs()
}

func newDeviceBindDeviceResult() interface{} {
	return device.NewDeviceBindDeviceResult()
}

func unbindDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*device.DeviceUnbindDeviceArgs)
	realResult := result.(*device.DeviceUnbindDeviceResult)
	success, err := handler.(device.Device).UnbindDevice(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDeviceUnbindDeviceArgs() interface{} {
	return device.NewDeviceUnbindDeviceArgs()
}

func newDeviceUnbindDeviceResult() interface{} {
	return device.NewDeviceUnbindDeviceResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetDeviceList(ctx context.Context, req *device.DeviceFilter) (r []*device.DeviceInfo, err error) {
	var _args device.DeviceGetDeviceListArgs
	_args.Req = req
	var _result device.DeviceGetDeviceListResult
	if err = p.c.Call(ctx, "GetDeviceList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetDeviceDetail(ctx context.Context, req *common.Req) (r *device.DeviceInfo, err error) {
	var _args device.DeviceGetDeviceDetailArgs
	_args.Req = req
	var _result device.DeviceGetDeviceDetailResult
	if err = p.c.Call(ctx, "GetDeviceDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateDeviceInfo(ctx context.Context, req *device.DeviceInfo) (r *device.DeviceInfo, err error) {
	var _args device.DeviceUpdateDeviceInfoArgs
	_args.Req = req
	var _result device.DeviceUpdateDeviceInfoResult
	if err = p.c.Call(ctx, "UpdateDeviceInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BindDevice(ctx context.Context, req *device.DeviceInfo) (r *device.DeviceInfo, err error) {
	var _args device.DeviceBindDeviceArgs
	_args.Req = req
	var _result device.DeviceBindDeviceResult
	if err = p.c.Call(ctx, "BindDevice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UnbindDevice(ctx context.Context, req *common.Req) (r *common.Empty, err error) {
	var _args device.DeviceUnbindDeviceArgs
	_args.Req = req
	var _result device.DeviceUnbindDeviceResult
	if err = p.c.Call(ctx, "UnbindDevice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}