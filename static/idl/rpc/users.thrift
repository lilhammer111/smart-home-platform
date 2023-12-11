namespace go users

struct AuthResp {
    1: required i32 ID;
    2: required i32 ExpiredAt;
    3: required string Token;
}

struct UserData {
    1: optional i8 Gender;
    2: optional i8 Age;
    3: optional i32 ID;
    4: optional string Mobile;
    5: optional string Username;
    6: optional string Profile;
}

struct CredentialsReq {
    1: required string SmsCode;
    2: required string Mobile;
    3: optional string Password;
}

service Users {
    UserData FindUser(1: i32 req);
    bool VerifyCredentials(1: CredentialsReq req);
    UserData CreateOrUpdateUser(1: UserData req);
    UserData DeleteUser(1: i32 req);
}