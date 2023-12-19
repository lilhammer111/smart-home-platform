namespace go micro_user
include "../http/common.thrift"
include "../http/user.thrift"
include "../http/auth.thrift"


struct RpcFreezeReq {
    1: optional string Username;
    2: optional string Mobile;
    3: optional string Email;
}

struct RpcFreezeResp {
    9: required bool   IsFrozen;
    10: optional string ThawedAt;
}


struct RpcSmsReq {
    1: required string Mobile;
}

struct RpcUserId {
    1: required i32 UserId;
}

struct RpcVerifyCodeReq {
    1: required string Mobile;
    2: required string SmsCode;
}

struct RpcVerifyUsernamePwdReq {
    1: required string Username;
    2: required string EntryPwd;
}

struct RpcVerifyEmailPwdReq {
    1: required string Email;
    2: required string EntryPwd;
}

struct RpcFindUserReq {
    1: required i32 UserId;
}

struct RpcFindUserByMobileReq {
    1: required string Mobile;
}

struct RpcFindUserByUsernameReq {
    1: required string Username;
}

struct RpcFindUserByOpenidReq {
    1: required string Openid;
}

service micro_user {
    common.Empty SendSmsViaAliyun(1: RpcSmsReq req);
    RpcFreezeResp FreezePatrolBeforeVerify(1: RpcFreezeReq req);
    RpcFreezeResp FreezePatrolAfterVerify(1: RpcFreezeReq req);
    common.Empty VerifySmsCode(1: RpcVerifyCodeReq req);
    common.Empty VerifyUsernamePwd(1: RpcVerifyUsernamePwdReq req);
    common.Empty VerifyEmailPwd(1: RpcVerifyEmailPwdReq req);
    user.UserInfo FindUser (1: RpcFindUserReq req);
    user.UserInfo FindUserByOpenid (1:RpcFindUserByOpenidReq req);
    user.UserInfo FindUserByMobile (1: RpcFindUserByMobileReq req);
    user.UserInfo FindUserByUsername (1: RpcFindUserByUsernameReq req);
    list<user.UserInfo> QueryUsersWithFilter(1: user.UsersFilter req);
    user.UserInfo UpdateUser(1: user.UserInfo req);
    user.UserInfo CreateUser(1: user.UserInfo req);
    common.Empty DeleteUser(1: i32 UserId);
}