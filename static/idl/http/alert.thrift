namespace go alert
include "common.thrift"

struct AlertFilter {
    2: optional i16 Page (api.query="page", api.vd="$>=0");
    1: optional i16 Limit (api.query="limit", api.vd="$>0");
    3: optional i8 Level (api.query="level", api.vd="in($, 0, 1, 2)");
    4: optional i32 DeviceId (api.query="device_id");
    5: optional string Sort (api.query="sort", api.vd="regexp('^[a-zA-Z]+_(asc|desc)$')");
    6: optional string StartDate (api.query="start_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
    7: optional string EndDate (api.query="end_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
}

struct AlertInfo {
    1: optional i32 Id (api.body="id");
    4: required i32 DeviceId (api.body="device_id");
    2: required i8 Count (api.body="count", api.vd="$<=127 && $>=0");
    3: required i8 Level (api.body="level", api.vd="in($, 0, 1, 2)");
    5: required string Desc (api.body="desc", api.vd="len($)<=600");
    6: required string FirstAlarm (api.body="first_alarm");
    7: required string LastAlarm (api.body="last_alarm");
}

service alert {
    list<AlertInfo> GetAlertList(1: AlertFilter req) (api.get="/api/devices/alerts/list");
    AlertInfo GetAlertDetail(1: common.Req req) (api.get="/api/devices/alerts/detail");
    AlertInfo UpdateAlertInfo(1: AlertInfo req) (api.put="/api/devices/alerts/update");
    AlertInfo UploadAlertInfo(1: AlertInfo req) (api.post="/api/devices/alerts/upload");
    common.Empty DeleteAlert(1: common.Req req) (api.delete="/api/devices/alerts/delete");
}