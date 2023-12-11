namespace go users
include "std_resp.thrift"

struct UserID {
    1: string ID (api.query="id");
}

struct UserInfo {
    1: required i32 ID (api.body="id");
    2: required i8 Age (api.body="age", api.vd="$>=0 && $ <=200>")
    3: required i8 Gender (api.body="gender", api.vd="in($,0,1,2)")
    4: required string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    5: required string Profile (api.body="profile", api.vd="regexp('^.{0,200}$')")
    6: required string Username (api.body="username", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')");
}

service users {
    std_resp.StdResp GetUserDetail(1: UserID req) (api.get="/api/users/detail");
    std_resp.StdResp UpdateUserInfo(1: UserInfo req) (api.put="/api/users/detail");
    std_resp.StdResp DeregisterUser(1: UserID req) (api.get="/api/users/detail");
}