namespace go users
include "std_resp.thrift"
include "std_req.thrift"

struct UsersFilter {
    1: optional i8 Gender (api.query="gender", api.vd="in($, 0, 1, 2)");
    2: optional i16 Page (api.query="page", api.vd="$>=0");
    3: optional i16 Limit (api.query="limit", api.vd="$>0");
    4: optional string Sort (api.query="sort", api.vd="regexp('^[a-zA-Z]+_(asc|desc)$')");
    5: optional string Search (api.query="search", api.vd="regexp('^[\p{L}\p{N}\p{Lo}]+$')");
    6: optional string StartDate (api.query="start_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
    7: optional string EndDate (api.query="end_date", api.vd="regexp('^\d{4}_\d{2}_\d{2}$')");
}

struct UserInfo {
    1: required i32 Id (api.body="id");
    2: required i8 Age (api.body="age", api.vd="$>=0 && $ <=200>");
    3: required i8 Gender (api.body="gender", api.vd="in($,0,1,2)");
    4: required string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    5: required string Profile (api.body="profile", api.vd="regexp('^.{0,200}$')");
    6: required string Username (api.body="username", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')"); //匹配中文字母数字下划线
}

service users {
    std_resp.StdResp GetUserList(1: UsersFilter req) (api.get="/api/users/list");
    std_resp.StdResp GetUserDetail(1: req.IdReq req) (api.get="/api/users/detail");
    std_resp.StdResp UpdateUserInfo(1: UserInfo req) (api.put="/api/users/update");
    std_resp.StdResp DeregisterUser(1: req.IdReq req) (api.delete="/api/users/deregister");
}