package micro_device_cli

import (
	"context"
	alert "git.zqbjj.top/pet/services/cmd/http/kitex_gen/alert"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	device "git.zqbjj.top/pet/services/cmd/http/kitex_gen/device"
	micro_device "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FindDevice(ctx context.Context, req *micro_device.RpcFindDeviceReq, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error) {
	resp, err = defaultClient.FindDevice(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindDevice call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryDevicesWithFilter(ctx context.Context, req *device.DeviceFilter, callOptions ...callopt.Option) (resp []*device.DeviceInfo, err error) {
	resp, err = defaultClient.QueryDevicesWithFilter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryDevicesWithFilter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error) {
	resp, err = defaultClient.CreateDevice(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateDevice call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateDevice(ctx context.Context, req *device.DeviceInfo, callOptions ...callopt.Option) (resp *device.DeviceInfo, err error) {
	resp, err = defaultClient.UpdateDevice(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateDevice call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteDevice(ctx context.Context, req *micro_device.RpcDeleteDeviceReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteDevice(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteDevice call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindLocationTitle(ctx context.Context, req *micro_device.RpcFindLocReq, callOptions ...callopt.Option) (resp *micro_device.LocationData, err error) {
	resp, err = defaultClient.FindLocationTitle(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindLocationTitle call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindAllLocationEnum(ctx context.Context, callOptions ...callopt.Option) (resp []*micro_device.LocationData, err error) {
	resp, err = defaultClient.FindAllLocationEnum(ctx, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindAllLocationEnum call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ReduceLocationEnum(ctx context.Context, req *micro_device.RpcReduceLocReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.ReduceLocationEnum(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ReduceLocationEnum call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ExpandLocationEnum(ctx context.Context, req *micro_device.RpcExpandLocReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.ExpandLocationEnum(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ExpandLocationEnum call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindAlert(ctx context.Context, req *micro_device.RpcFindAlertReq, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error) {
	resp, err = defaultClient.FindAlert(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindAlert call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryAlertsWithFilter(ctx context.Context, req *alert.AlertFilter, callOptions ...callopt.Option) (resp []*alert.AlertInfo, err error) {
	resp, err = defaultClient.QueryAlertsWithFilter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryAlertsWithFilter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error) {
	resp, err = defaultClient.CreateAlert(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateAlert call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateAlert(ctx context.Context, req *alert.AlertInfo, callOptions ...callopt.Option) (resp *alert.AlertInfo, err error) {
	resp, err = defaultClient.UpdateAlert(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateAlert call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteAlert(ctx context.Context, req *micro_device.RpcDeleteAlertReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteAlert(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteAlert call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
