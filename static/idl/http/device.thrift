namespace go device
include "common.thrift"


struct DeviceFilter {
    1: optional i8 State (api.query="state")
    2: optional i16 Page (api.query="page");
    3: optional i16 Limit (api.query="limit");
    4: optional list<string> Sorts (api.query="sorts");
    5: optional string Search (api.query="search");
    6: optional string StartDate (api.query="start_date");
    7: optional string EndDate (api.query="end_date");
}

struct DeviceInfo {
    1: optional i32 Id (api.body="id");
    10: required i32 OwnerId (api.body="owner_id");
    11: required i32 ProductId (api.body="product_id");
    2: required string SerialNo (api.body="serial_no");
    3: required string Name (api.body="name", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')");
    4: required i8 State (api.body="state")
    6: required i8 LocationId (api.body="location" );
    7: required string HardwareVersion (api.body="hardware_version");
    8: required string SoftwareVersion (api.body="software_version");
    9: required string Desc (api.body="desc");
}

service device {
    list<DeviceInfo> GetDeviceList(1: DeviceFilter req) (api.get="/api/devices/list");
    DeviceInfo GetDeviceDetail(1: common.Req req) (api.get="/api/devices/detail");
    DeviceInfo UpdateDeviceInfo(1: DeviceInfo req) (api.put="/api/devices/update");
    DeviceInfo BindDevice(1: DeviceInfo req) (api.post="/api/devices/bind");
    common.Empty UnbindDevice(1: common.Req req) (api.delete="/api/devices/unbind");
}