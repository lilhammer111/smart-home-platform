namespace go micro_user
include "../http/user.thrift"
include "common_rpc.thrift"

struct CredentialRpcReq {
    1: required string SmsCode;
    2: required string Mobile;
    3: optional string Password;
    4: optional string Username;
}

struct FreezeRpcReq {
    1: optional string Username;
    2: optional string Mobile;
    3: optional string Email;
}

struct FreezeRpcResp {
    9: required bool   IsFrozen (api.body="is_frozen");
    10: optional string ThawedAt (api.body="thawed_at");
}




service micro_user {
    FreezeRpcResp FreezePatrolBeforeAuth(1: FreezeRpcReq req);
    FreezeRpcResp FreezePatrolAfterAuth(1: common_rpc.IdRpcReq req);
    common_rpc.EmptyRpcResp VerifyCredentials(1: CredentialRpcReq req);

    user.UserInfo FindUser (1: common_rpc.IdRpcReq req);
    list<user.UserInfo> QueryUsersWithFilter(1: user.UsersFilter req)
    user.UserInfo UpsertUser(1: user.UserInfo req);
    common_rpc.EmptyRpcResp DeleteUser(1: common_rpc.IdRpcReq req);
}