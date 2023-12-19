namespace go micro_user
include "../http/common.thrift"
include "../http/user.thrift"
include "../http/auth.thrift"


struct RpcFreezeReq {
    1: optional string Username;
    2: optional string Mobile;
    3: optional string Email;
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

struct RpcVerifyResp {
    1: required bool VerifyPassed;
}

struct RpcAfterVerifyReq {
    1: required i32 UserId;
    2: required bool VerifyPassed;
}

struct RpcRequestOpenIdResp {
    1: required string OpenId;
}

struct RpcRequestOpenIdReq {
    1: required string JsCode;
    2: required string Secret;
    3: required string Appid;
    4: optional string Unionid;
}

struct RpcDeleteUserReq {
    1: required i32 UserId;
}

service micro_user {
    common.Empty SendSmsViaAliyun(1: RpcSmsReq req);
    RpcUserId FreezePatrolBeforeVerify(1: RpcFreezeReq req);
    common.Empty FreezePatrolAfterVerify(1: RpcAfterVerifyReq req);
    RpcVerifyResp VerifySmsCode(1: RpcVerifyCodeReq req);
    RpcVerifyResp VerifyUsernamePwd(1: RpcVerifyUsernamePwdReq req);
    RpcVerifyResp VerifyEmailPwd(1: RpcVerifyEmailPwdReq req);
    user.UserInfo FindUser (1: RpcFindUserReq req);
    user.UserInfo FindUserByOpenid (1:RpcFindUserByOpenidReq req);
    user.UserInfo FindUserByMobile (1: RpcFindUserByMobileReq req);
    user.UserInfo FindUserByUsername (1: RpcFindUserByUsernameReq req);
    list<user.UserInfo> QueryUsersWithFilter(1: user.UsersFilter req);
    user.UserInfo UpdateUser(1: user.UserInfo req);
    user.UserInfo CreateUser(1: user.UserInfo req);
    common.Empty DeleteUser(1: RpcDeleteUserReq req);
    RpcRequestOpenIdResp RequestOpenId(1: RpcRequestOpenIdReq req);
}