namespace go micro_user
include "../common.thrift"
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

struct RpcOpenIdReq {
    1: required string OpenId;
}



service micro_user {
    common.Empty SendSmsViaAliyun(1: RpcSmsReq req);
    RpcFreezeResp FreezePatrolBeforeAuth(1: RpcFreezeReq req);
    RpcFreezeResp FreezePatrolAfterAuth(1: i32 UserId);
    common.Empty VerifySmsCode(1: string mobile; 2: string smsCode);
    common.Empty VerifyUsernamePwd(1: string username; 2: string entryPwd);
    common.Empty VerifyEmailPwd(1: string email; 2: string entryPwd);

    user.UserInfo FindUser (1: i32 UserId);
    list<user.UserInfo> QueryUsersWithFilter(1: RpcUsersFilterReq req)
    user.UserInfo UpdateUser(1: user.UserInfo req);
    user.UserInfo CreateUser(1: user.UserInfo req);
    common.Empty DeleteUser(1: i32 UserId);
}