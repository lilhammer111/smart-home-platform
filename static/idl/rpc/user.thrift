namespace go user

struct AuthResp {
    1: required i32 Id;
    2: required i32 ExpiredAt;
    3: required string Token;
}

struct UserData {
    1: optional i8 Gender (go.tag='json:"gender"');
    2: required i8 Age;
    3: required i32 Id;
    4: required string Mobile;
    5: required string Username;
    6: required string Profile;
}

struct CredentialsReq {
    1: required string SmsCode;
    2: required string Mobile;
    3: optional string Password;
    4: optional string Username;
}

struct UsersFilterReq {
    1: optional i8 Page;
}

service user {
    UserData FindUserByID(1: i32 req);
    list<UserData> GetUsersByFilter(1: UsersFilterReq req)
    bool VerifyCredentials(1: CredentialsReq req);
    UserData CreateOrUpdateUser(1: UserData req);
    UserData DeleteUser(1: i32 req);
}