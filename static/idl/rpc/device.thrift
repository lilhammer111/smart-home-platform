namespace go micro_device
include "../http/device.thrift"
include "../http/alert.thrift"
include "../http/common.thrift"

struct RpcExpandLocReq {
    1: required string Title;
}

struct RpcReduceLocReq {
    1: required i16 Id;
}

struct RpcFindLocReq {
    1: required i16 Id;
}

struct LocationData {
    1: required i16 Id;
    2: required string Title;
}

struct LocationListResp {
    1: required list<LocationData> LocationList;
}

struct RpcFindDeviceReq {
    1: required i32 Id;
}

struct RpcDeleteDeviceReq {
    1: required i32 Id;
}

struct RpcFindAlertReq {
    1: required i32 Id;
}

struct RpcDeleteAlertReq {
    1: required i32 Id;
}


service MicroDevice {
    device.DeviceInfo FindDevice(1: RpcFindDeviceReq req);
    list<device.DeviceInfo> QueryDevicesWithFilter(1: device.DeviceFilter req);
    device.DeviceInfo CreateDevice(1: device.DeviceInfo req);
    device.DeviceInfo UpdateDevice(1: device.DeviceInfo req);
    common.Empty DeleteDevice(1: RpcDeleteDeviceReq req)

    LocationData FindLocationTitle(1: RpcFindLocReq req);
    list<LocationData> FindAllLocationEnum();
    common.Empty ReduceLocationEnum(1: RpcReduceLocReq req);
    common.Empty ExpandLocationEnum(1: RpcExpandLocReq req);

    alert.AlertInfo FindAlert(1: RpcFindAlertReq req);
    list<alert.AlertInfo> QueryAlertsWithFilter(1: alert.AlertFilter req);
    alert.AlertInfo CreateAlert(1: alert.AlertInfo req);
    alert.AlertInfo UpdateAlert(1: alert.AlertInfo req);
    common.Empty DeleteAlert(1: RpcDeleteAlertReq req)
}

service DeviceService {

}
