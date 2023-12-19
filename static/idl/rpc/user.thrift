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

struct RpcUsersFilterReq {
    1: optional i8 Gender;
    2: optional i16 Page;
    3: optional i16 Limit;
    4: optional string Sort;
    5: optional string Search;
    6: optional string StartDate;
    7: optional string EndDate;
}

//struct RpcUser {
//    1: optional i32 Id;
//    2: required i8 Age;
//    3: required i8 Gender;
//    4: required string Mobile;
//    5: required string Profile;
//    6: required string Username; //匹配中文字母数字下划线
//    7: required string Email;
//    8: required string Avatar;
//}

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
    list<user.UserInfo> QueryUsersWithFilter(1: RpcUsersFilterReq req);
    user.UserInfo UpdateUser(1: user.UserInfo req);
    user.UserInfo CreateUser(1: user.UserInfo req);
    common.Empty DeleteUser(1: i32 UserId);
}