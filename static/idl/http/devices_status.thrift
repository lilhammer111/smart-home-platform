namespace go devices_status
include "std_resp.thrift"
include "std_req.thrift"

struct RecentFeeding {
    1: required i8 FeedingAmount (api.body="feeding_amount");
    2: required i32 timestamp(api.body="id"); // until 2038.1.19
    3: required string FoodType (api.body="food_type");
}

struct DeviceStatusInfo {
    1: required bool IsOnline (api.body="is_online");
    2: required i8 DeviceStatus (api.body="device_status");
    3: required i8 FoodSurplus (api.body="food_surplus");
    4: required i8 BatteryLevel (api.body="battery_level", api.vd="$<=100 && $>=0");
    5: optional i32 Id (api.body="id");
    6: required i32 DeviceId (api.body="device_id");
    7: required RecentFeeding recent_feeding (api.body="recent_feeding");
}

service devices_status {
    std_resp.StdResp GetDeviceStatus(1: req.IdReq req) (api.get="/api/devices/status");
    std_resp.StdResp InitDeviceStatus(1: DeviceStatusInfo req) (api.post="/api/devices/status");
    std_resp.StdResp UpdateDeviceStatus(1: DeviceStatusInfo req) (api.put="/api/devices/status");
}