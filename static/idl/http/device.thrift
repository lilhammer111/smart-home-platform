namespace go device
include "common_http.thrift"


struct DeviceFilter {
    1: optional bool IsOnline (api.query="is_online");
    2: optional i16 Page (api.query="page", api.vd="$>=0");
    3: optional i16 Limit (api.query="limit", api.vd="$>0");
    4: optional string Sort (api.query="sort", api.vd="regexp('^[a-zA-Z]+_(asc|desc)$')");
    5: optional string Search (api.query="search", api.vd="regexp('^[\p{L}\p{N}\p{Lo}]+$')");
    6: optional string StartDate (api.query="start_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
    7: optional string EndDate (api.query="end_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
}

struct DeviceInfo {
    1: optional i32 Id (api.body="id");
    2: required string SerialNo (api.body="serial_no");
    3: required string Name (api.body="name", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')");
    4: required string Model (api.body="model");
    5: required string Type (api.body="type");
    6: required string Location (api.body="location");
    7: required string HardwareVersion (api.body="hardware_version");
    8: required string SoftwareVersion (api.body="software_version");
    9: required string Desc (api.body="desc");
}

// response
struct DeviceListResp {
    1: required common_http.Resp Meta;
    2: required list<i32> Data;
}

struct DeviceInfoResp {
    1: required common_http.Resp Meta;
    2: required DeviceInfo Data;
}


service device {
    DeviceListResp GetDeviceList(1: DeviceFilter req) (api.get="/api/devices/list");
    DeviceInfoResp GetDeviceDetail(1: common_http.Req req) (api.get="/api/devices/detail");
    DeviceInfoResp UpdateDeviceInfo(1: DeviceInfo req) (api.put="/api/devices/update");
    DeviceInfoResp BindDevice(1: DeviceInfo req) (api.post="/api/devices/bind");
    common_http.Resp UnbindDevice(1: common_http.Req req) (api.delete="/api/devices/unbind");
}