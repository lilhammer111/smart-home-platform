namespace go device
include "common.thrift"


struct DeviceFilter {
    1: optional i8 State (api.query="state", api.vd="len($)<=8 && len($)>=0" ,go.tag="example:\"7\"")
    2: optional i16 Page (api.query="page", api.vd="$>=0" ,go.tag="example:\"1\"");
    3: optional i16 Limit (api.query="limit", api.vd="$>0" ,go.tag="example:\"1\"");
    4: optional string Sort (api.query="sort", api.vd="regexp('^[a-zA-Z]+_(asc|desc)$')" ,go.tag="example:\"name_asc\"");
    5: optional string Search (api.query="search", api.vd="regexp('^[\p{L}\p{N}\p{Lo}]+$')" ,go.tag="example:\"hello,device\"");
    6: optional string StartDate (api.query="start_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')" ,go.tag="example:\"2023_02_01\"");
    7: optional string EndDate (api.query="end_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')" ,go.tag="example:\"2023_02_05\"");
}

struct DeviceInfo {
    1: optional i32 Id (api.body="id",go.tag="example:\"1\"");
    10: required i32 OwnerId (api.body="id" ,go.tag="example:\"1\"");
    11: required i32 ProductId (api.body="id" ,go.tag="example:\"1\"");
    2: required string SerialNo (api.body="serial_no" ,go.tag="example:\"KEwju0rKOlKCDAxUnDzQI\"");
    3: required string Name (api.body="name", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')" ,go.tag="example:\"demon's feeder\"");
    4: required i8 State (api.body="state", go.tag="example:\"8\"")
    6: required string LocationId (api.body="location" ,go.tag="example:\"1\"");
    7: required string HardwareVersion (api.body="hardware_version" ,go.tag="example:\"v 1.0.0\"");
    8: required string SoftwareVersion (api.body="software_version" ,go.tag="example:\"v 1.0.0\"");
    9: required string Desc (api.body="desc" ,go.tag="example:\"hello,device\"");
}

service device {
    list<DeviceInfo> GetDeviceList(1: DeviceFilter req) (api.get="/api/devices/list");
    DeviceInfo GetDeviceDetail(1: common.Req req) (api.get="/api/devices/detail");
    DeviceInfo UpdateDeviceInfo(1: DeviceInfo req) (api.put="/api/devices/update");
    DeviceInfo BindDevice(1: DeviceInfo req) (api.post="/api/devices/bind");
    common.Empty UnbindDevice(1: common.Req req) (api.delete="/api/devices/unbind");
}