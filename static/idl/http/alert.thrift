namespace go alert
include "common.thrift"

struct AlertFilter {
    2: optional i16 Page (api.query="page", api.vd="$>=0" ,go.tag="example:\"1\"");
    1: optional i16 Limit (api.query="limit", api.vd="$>0" ,go.tag="example:\"1\"");
    3: optional i8 Level (api.query="level", api.vd="in($, 0, 1, 2)" ,go.tag="example:\"1\"");
    4: optional i32 DeviceId (api.query="device_id" ,go.tag="example:\"1\"");
    5: optional list<string> Sorts (api.query="sorts", api.vd="range($, regexp('^[a-zA-Z]+ (asc|desc)$', #v))" ,go.tag="example:\"device_id desc\"");
    6: optional string StartDate (api.query="start_date", api.vd="regexp('^\d{4}-\d{2}-\d{2}$')" ,go.tag="example:\"2023-12-22\"");
    7: optional string EndDate (api.query="end_date", api.vd="regexp('^\d{4}-\d{2}-\d{2}$')" ,go.tag="example:\"2024-02-10\"");
    8: optional bool IsOngoing (api.query="is_ongoing" ,go.tag="example:\"true\"");
}

struct AlertInfo {
    1: optional i32 Id (api.body="id" ,go.tag="example:\"1\"");
    4: required i32 DeviceId (api.body="device_id" ,go.tag="example:\"1\"");
    2: required i8 Count (api.body="count", api.vd="$<=127 && $>=0" ,go.tag="example:\"50\"");
    3: required i8 Level (api.body="level", api.vd="in($, 0, 1, 2)" ,go.tag="example:\"1\"");
    5: required string Desc (api.body="desc", api.vd="len($)<=600" ,go.tag="example:\"hello,alert\"");
    6: required string FirstAlarm (api.body="first_alarm" ,go.tag="example:\"2023-12-22 15:16:00\"");
    7: required string LastAlarm (api.body="last_alarm" ,go.tag="example:\"2023-12-30 15:16:00\"");
    8: required bool IsOngoing (api.body="is_ongoing" ,go.tag="example:\"true\"");
}

service alert {
    list<AlertInfo> GetAlertList(1: AlertFilter req) (api.get="/api/devices/alerts/list");
    AlertInfo GetAlertDetail(1: common.Req req) (api.get="/api/devices/alerts/detail");
    AlertInfo UpdateAlertInfo(1: AlertInfo req) (api.put="/api/devices/alerts/update");
    AlertInfo UploadAlertInfo(1: AlertInfo req) (api.post="/api/devices/alerts/upload");
    common.Empty DeleteAlert(1: common.Req req) (api.delete="/api/devices/alerts/delete");
}