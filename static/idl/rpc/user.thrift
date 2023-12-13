namespace go micro_user
include "common_rpc.thrift"

struct RpcCredentialReq {
    1: required string SmsCode;
    2: required string Mobile;
    3: optional string Password;
    4: optional string Username;
}

struct RpcFreezeReq {
    1: optional string Username;
    2: optional string Mobile;
    3: optional string Email;
}

struct RpcFreezeResp {
    9: required bool   IsFrozen (api.body="is_frozen");
    10: optional string ThawedAt (api.body="thawed_at");
}

struct RpcUsersFilterReq {
    1: optional i8 Gender;
    2: optional i16 Page;
    3: optional i16 Limit;
    4: optional string Sort;
    5: optional string Search;
    6: optional string StartDate;
    7: optional string EndDate;
}

struct RpcUser {
    1: required i32 Id;
    2: required i8 Age;
    3: required i8 Gender;
    4: required string Mobile;
    5: required string Profile;
    6: required string Username; //匹配中文字母数字下划线
    7: required string Email;
    8: required string Avatar;
}


service micro_user {
    RpcFreezeResp FreezePatrolBeforeAuth(1: RpcFreezeReq req);
    RpcFreezeResp FreezePatrolAfterAuth(1: common_rpc.RpcId req);
    common_rpc.RpcEmpty VerifyCredentials(1: RpcCredentialReq req);

    RpcUser FindUser (1: common_rpc.RpcId req);
    list<RpcUser> QueryUsersWithFilter(1: RpcUsersFilterReq req)
    RpcUser UpsertUser(1: RpcUser req);
    common_rpc.RpcEmpty DeleteUser(1: common_rpc.RpcId req);
}