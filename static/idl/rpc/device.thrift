namespace go device_micro
include "../http/device.thrift"
include "../http/device_status.thrift"
include "../http/alert.thrift"
include "common_rpc.thrift"

struct LocationReq {
    1: required string Name;
}

struct LocationListResp {
    1: required list<string> LocationList;
}

service Device {
    device.DeviceInfo FindDevice(1: common_rpc.IdRpcReq req);
    list<device.DeviceInfo> QueryDevicesWithFilter(1: device.DeviceFilter req);
    device.DeviceInfo UpsertDevice(1: device.DeviceInfo req);
    void DeleteDevice(1: common_rpc.IdRpcReq req)
}

service LocationEnum {
    LocationListResp FindLocationEnum()
    LocationListResp ReduceLocationEnum(1: LocationReq req)
    LocationListResp ExpandLocationEnum(1: LocationReq req)
}

service DeviceStatus {
    Query
}

service DeviceAlert {

}
