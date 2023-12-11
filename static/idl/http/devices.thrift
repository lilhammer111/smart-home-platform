namespace go devices
include "std_resp.thrift"
include "std_req.thrift"

struct DevicesFilter {
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
    2: required string DeviceSerialNo (api.body="device_serial_no");
    3: required string DeviceName (api.body="device_name", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')");
    4: required string DeviceModel (api.body="device_model");
    5: required string DeviceType (api.body="device_type");
    6: required string Location (api.body="location");
    7: required string HardwareVersion (api.body="hardware_version");
    8: required string SoftwareVersion (api.body="software_version");
}

service devices {
    std_std_resp.StdResp GetDeviceList(1: DevicesFilter req) (api.get="/api/devices");
    std_resp.StdResp GetDeviceDetail(1: req.IdReq req) (api.get="/api/devices/detail");
    std_resp.StdResp UpdateDeviceInfo(1: DeviceInfo req) (api.put="/api/devices");
    std_resp.StdResp BindDevice(1: DeviceInfo req) (api.post="/api/devices");
    std_resp.StdResp UnbindDevice(1: req.IdReq req) (api.delete="/api/devices");
}