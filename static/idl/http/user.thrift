namespace go user
include "common.thrift"

struct UsersFilter {
    1: optional i16 Page (api.query="page");
    2: optional i16 Limit (api.query="limit");
}

struct UserInfo {
    1: optional i32 Id (api.body="id");
    2: required i8 Age (api.body="age", api.vd="$>=0 && $<=200");
    3: required i8 Gender (api.body="gender", api.vd="in($,0,1,2)");
    4: required string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    5: required string Profile (api.body="profile", api.vd="regexp('^.{0,200}$')");
    6: required string Username (api.body="username", api.vd="regexp('^[a-z0-9_]{1,30}$')"); // Only lowercase letters, numbers and underscores are allowed
    7: required string Password (api.body="password", api.vd="regexp('^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{6,12}$
')");
    8: required string Email (api.body="email", api.vd="regexp('^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$')");
    9: required string Avatar (api.body="avatar", api.vd="regexp('^.{0,200}$')");
    10: required string Openid (api.body="openid", api.vd="len($)==28");
}

struct UserInfoResp {
    1: required i32 Id (api.body="id");
    2: required i8 Age (api.body="age,omitempty", api.vd="$>=0 && $<=200");
    3: required i8 Gender (api.body="gender,omitempty", api.vd="in($,0,1,2)");
    4: required string Mobile (api.body="mobile,omitempty", api.vd="regexp('^1[3-9]\\d{9}$')");
    5: required string Profile (api.body="profile,omitempty", api.vd="regexp('^.{0,200}$')");
    6: required string Username (api.body="username,omitempty", api.vd="regexp('^[a-z0-9_]{1,30}$')"); // Only lowercase letters, numbers and underscores are allowed
    8: required string Email (api.body="email,omitempty", api.vd="regexp('^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$')");
    9: required string Avatar (api.body="avatar,omitempty", api.vd="regexp('^.{0,200}$')");
}

service user {
    UserInfoResp GetCurUserInfo(1: common.Empty req) (api.get="/api/users/current");
    list<UserInfoResp> GetUserList(1: UsersFilter req) (api.get="/api/users/list");
    UserInfoResp GetUserDetail(1: common.Req req) (api.get="/api/users/detail");
    UserInfoResp UpdateUserInfo(1: UserInfo req) (api.put="/api/users/update");
    common.Empty DeregisterUser(1: common.Req req) (api.delete="/api/users/deregister");
}