package micro_device_cli

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/alert"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/device"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_device"

	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_device/microdevice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() microdevice.Client
	Service() string
	FindDevice(ctx context.Context, req *micro_device.RpcFindDeviceReq, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error)

	QueryDevicesWithFilter(ctx context.Context, req *device.DeviceFilter, callOptions ...callopt.Option) (resp []*device.DeviceInfo, err error)

	CreateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error)

	UpdateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error)

	DeleteDevice(ctx context.Context, req *micro_device.RpcDeleteDeviceReq, callOptions ...callopt.Option) (resp *common.Empty, err error)

	FindLocationTitle(ctx context.Context, req *micro_device.RpcFindLocReq, callOptions ...callopt.Option) (resp *micro_device.LocationData, err error)

	FindAllLocationEnum(ctx context.Context, callOptions ...callopt.Option) (resp []*micro_device.LocationData, err error)

	ReduceLocationEnum(ctx context.Context, req *micro_device.RpcReduceLocReq, callOptions ...callopt.Option) (resp *common.Empty, err error)

	ExpandLocationEnum(ctx context.Context, req *micro_device.RpcExpandLocReq, callOptions ...callopt.Option) (resp *common.Empty, err error)

	FindAlert(ctx context.Context, req *micro_device.RpcFindAlertReq, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error)

	QueryAlertsWithFilter(ctx context.Context, req *alert.AlertFilter, callOptions ...callopt.Option) (resp []*alert.AlertInfo, err error)

	CreateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error)

	UpdateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error)

	DeleteAlert(ctx context.Context, req *micro_device.RpcDeleteAlertReq, callOptions ...callopt.Option) (resp *common.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := microdevice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient microdevice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() microdevice.Client {
	return c.kitexClient
}

func (c *clientImpl) FindDevice(ctx context.Context, req *micro_device.RpcFindDeviceReq, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error) {
	return c.kitexClient.FindDevice(ctx, req, callOptions...)
}

func (c *clientImpl) QueryDevicesWithFilter(ctx context.Context, req *device.DeviceFilter, callOptions ...callopt.Option) (resp []*device.DeviceInfo, err error) {
	return c.kitexClient.QueryDevicesWithFilter(ctx, req, callOptions...)
}

func (c *clientImpl) CreateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error) {
	return c.kitexClient.CreateDevice(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error) {
	return c.kitexClient.UpdateDevice(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteDevice(ctx context.Context, req *micro_device.RpcDeleteDeviceReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteDevice(ctx, req, callOptions...)
}

func (c *clientImpl) FindLocationTitle(ctx context.Context, req *micro_device.RpcFindLocReq, callOptions ...callopt.Option) (resp *micro_device.LocationData, err error) {
	return c.kitexClient.FindLocationTitle(ctx, req, callOptions...)
}

func (c *clientImpl) FindAllLocationEnum(ctx context.Context, callOptions ...callopt.Option) (resp []*micro_device.LocationData, err error) {
	return c.kitexClient.FindAllLocationEnum(ctx, callOptions...)
}

func (c *clientImpl) ReduceLocationEnum(ctx context.Context, req *micro_device.RpcReduceLocReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.ReduceLocationEnum(ctx, req, callOptions...)
}

func (c *clientImpl) ExpandLocationEnum(ctx context.Context, req *micro_device.RpcExpandLocReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.ExpandLocationEnum(ctx, req, callOptions...)
}

func (c *clientImpl) FindAlert(ctx context.Context, req *micro_device.RpcFindAlertReq, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error) {
	return c.kitexClient.FindAlert(ctx, req, callOptions...)
}

func (c *clientImpl) QueryAlertsWithFilter(ctx context.Context, req *alert.AlertFilter, callOptions ...callopt.Option) (resp []*alert.AlertInfo, err error) {
	return c.kitexClient.QueryAlertsWithFilter(ctx, req, callOptions...)
}

func (c *clientImpl) CreateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error) {
	return c.kitexClient.CreateAlert(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error) {
	return c.kitexClient.UpdateAlert(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteAlert(ctx context.Context, req *micro_device.RpcDeleteAlertReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteAlert(ctx, req, callOptions...)
}
