namespace go device_status
include "standard.thrift"

struct RecentFeeding {
    1: required i8 Amount (api.body="feeding_amount");
    2: required i32 Timestamp(api.body="timestamp"); // until 2038.1.19
    3: required string FoodType (api.body="food_type");
}

struct DeviceStatusInfo {
    1: optional i32 Id (api.body="id");
    2: required i32 DeviceId (api.body="device_id");
    3: required bool IsFaulty (api.body="is_faulty");
    4: required bool PoweredOn (api.body="powered_on");
    5: required bool IsConnected (api.body="is_connected");
    6: required bool IsWorking (api.body="is_working");
    7: required i8 FoodSurplus (api.body="food_surplus");
    8: required i8 BatteryCharge (api.body="battery_charge", api.vd="$<=100 && $>=0");
    9: required RecentFeeding RecentFeeding (api.body="recent_feeding");
}

service device_status {
    standard.Resp GetDeviceStatus(1: i32 req) (api.get="/api/devices/status/get");
    standard.Resp InitDeviceStatus(1: DeviceStatusInfo req) (api.post="/api/devices/status/init");
    standard.Resp UpdateDeviceStatus(1: DeviceStatusInfo req) (api.put="/api/devices/status/update");
}