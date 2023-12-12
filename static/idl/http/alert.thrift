namespace go alert
include "standard.thrift"

struct AlertFilter {
    1: optional i16 Limit (api.query="limit", api.vd="$>0");
    2: optional i16 Page (api.query="page", api.vd="$>=0");
    3: optional i8 AlertLevel (api.query="alert_level", api.vd="in($, 0, 1, 2)");
    4: optional i32 DeviceId (api.query="device_id");
    5: optional string Sort (api.query="sort", api.vd="regexp('^[a-zA-Z]+_(asc|desc)$')");
    6: optional string StartDate (api.query="start_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
    7: optional string EndDate (api.query="end_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
}

struct AlertInfo {
    1: optional i32 Id (api.body="id");
    2: required i8 alert_times (api.body="alert_times", api.vd="$<=127 && $>=0");
    3: required i8 AlertLevel (api.body="alert_level", api.vd="in($, 0, 1, 2)");
    4: required i32 DeviceId (api.body="device_id");
    5: required string alert_desc (api.body="alert_desc", api.vd="len($)<=600");
}

// response
struct AlertListResp {
    1: required standard.Resp Meta;
    2: required list<i32> Data;
}

struct AlertInfoResp {
    1: required standard.Resp Meta;
    2: required AlertInfo Data;
}


service alert {
    AlertListResp GetAlertList(1: AlertFilter req) (api.get="/api/devices/alerts/list");
    AlertInfoResp GetAlertDetail(1: i32 req) (api.get="/api/devices/alerts/detail");
    AlertInfoResp UpdateAlertInfo(1: AlertInfo req) (api.put="/api/devices/alerts/update");
    AlertInfoResp UploadAlertInfo(1: AlertInfo req) (api.post="/api/devices/alerts/upload");
    standard.Resp DeleteAlert(1: i32 req) (api.delete="/api/devices/alerts/delete");
}