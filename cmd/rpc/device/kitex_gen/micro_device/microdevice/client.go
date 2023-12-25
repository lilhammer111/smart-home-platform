// Code generated by Kitex v0.6.1. DO NOT EDIT.

package microdevice

import (
	"context"
	alert "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/alert"
	common "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/common"
	device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
	micro_device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FindDevice(ctx context.Context, req *micro_device.RpcFindDeviceReq, callOptions ...callopt.Option) (r *device.DeviceInfo, err error)
	QueryDevicesWithFilter(ctx context.Context, req *device.DeviceFilter, callOptions ...callopt.Option) (r []*device.DeviceInfo, err error)
	CreateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (r *device.DeviceInfo, err error)
	UpdateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (r *device.DeviceInfo, err error)
	DeleteDevice(ctx context.Context, req *micro_device.RpcDeleteDeviceReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	FindLocationTitle(ctx context.Context, req *micro_device.RpcFindLocReq, callOptions ...callopt.Option) (r *micro_device.LocationData, err error)
	FindAllLocationEnum(ctx context.Context, callOptions ...callopt.Option) (r []*micro_device.LocationData, err error)
	ReduceLocationEnum(ctx context.Context, req *micro_device.RpcReduceLocReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	ExpandLocationEnum(ctx context.Context, req *micro_device.RpcExpandLocReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	FindAlert(ctx context.Context, req *micro_device.RpcFindAlertReq, callOptions ...callopt.Option) (r *alert.AlertInfo, err error)
	QueryAlertsWithFilter(ctx context.Context, req *alert.AlertFilter, callOptions ...callopt.Option) (r []*alert.AlertInfo, err error)
	CreateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (r *alert.AlertInfo, err error)
	UpdateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (r *alert.AlertInfo, err error)
	DeleteAlert(ctx context.Context, req *micro_device.RpcDeleteAlertReq, callOptions ...callopt.Option) (r *common.Empty, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMicroDeviceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMicroDeviceClient struct {
	*kClient
}

func (p *kMicroDeviceClient) FindDevice(ctx context.Context, req *micro_device.RpcFindDeviceReq, callOptions ...callopt.Option) (r *device.DeviceInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindDevice(ctx, req)
}

func (p *kMicroDeviceClient) QueryDevicesWithFilter(ctx context.Context, req *device.DeviceFilter, callOptions ...callopt.Option) (r []*device.DeviceInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryDevicesWithFilter(ctx, req)
}

func (p *kMicroDeviceClient) CreateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (r *device.DeviceInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateDevice(ctx, req)
}

func (p *kMicroDeviceClient) UpdateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (r *device.DeviceInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateDevice(ctx, req)
}

func (p *kMicroDeviceClient) DeleteDevice(ctx context.Context, req *micro_device.RpcDeleteDeviceReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteDevice(ctx, req)
}

func (p *kMicroDeviceClient) FindLocationTitle(ctx context.Context, req *micro_device.RpcFindLocReq, callOptions ...callopt.Option) (r *micro_device.LocationData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindLocationTitle(ctx, req)
}

func (p *kMicroDeviceClient) FindAllLocationEnum(ctx context.Context, callOptions ...callopt.Option) (r []*micro_device.LocationData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindAllLocationEnum(ctx)
}

func (p *kMicroDeviceClient) ReduceLocationEnum(ctx context.Context, req *micro_device.RpcReduceLocReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReduceLocationEnum(ctx, req)
}

func (p *kMicroDeviceClient) ExpandLocationEnum(ctx context.Context, req *micro_device.RpcExpandLocReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ExpandLocationEnum(ctx, req)
}

func (p *kMicroDeviceClient) FindAlert(ctx context.Context, req *micro_device.RpcFindAlertReq, callOptions ...callopt.Option) (r *alert.AlertInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindAlert(ctx, req)
}

func (p *kMicroDeviceClient) QueryAlertsWithFilter(ctx context.Context, req *alert.AlertFilter, callOptions ...callopt.Option) (r []*alert.AlertInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryAlertsWithFilter(ctx, req)
}

func (p *kMicroDeviceClient) CreateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (r *alert.AlertInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateAlert(ctx, req)
}

func (p *kMicroDeviceClient) UpdateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (r *alert.AlertInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateAlert(ctx, req)
}

func (p *kMicroDeviceClient) DeleteAlert(ctx context.Context, req *micro_device.RpcDeleteAlertReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteAlert(ctx, req)
}
