namespace go user
include "common.thrift"

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
    1: optional i32 Id (api.body="id");
    2: optional i8 Age (api.body="age", api.vd="$>=0 && $ <=200>");
    3: optional i8 Gender (api.body="gender", api.vd="in($,0,1,2)");
    4: optional string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    5: optional string Profile (api.body="profile", api.vd="regexp('^.{0,200}$')");
    6: optional string Username (api.body="username", api.vd="regexp('^[a-z0-9_]{1,30}$')"); // Only lowercase letters, numbers and underscores are allowed
    7: optional string Email (api.body="email", api.vd="regexp('^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$')");
    8: optional string Avatar (api.body="avatar", api.vd="regexp('^.{0,200}$')");
}


service user {
    list<UserInfo> GetUserList(1: UsersFilter req) (api.get="/api/users/list");
    list<UserInfo> GetUserDetail(1: common.Req req) (api.get="/api/users/detail");
    list<UserInfo> UpdateUserInfo(1: UserInfo req) (api.put="/api/users/update");
    common.Empty DeregisterUser(1: common.Req req) (api.delete="/api/users/deregister");
}