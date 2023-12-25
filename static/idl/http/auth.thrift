namespace go auth
include "common.thrift"

struct SendSmsReq {
    1: required string Mobile (api.query="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
}

struct MobileRegisterReq {
    1: required string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    2: required string SmsCode (api.body="sms_code", api.vd="regexp('^\\d{6}$')");
    3: required string Username (api.body="username", api.vd="regexp('^[a-z0-9_]{1,30}$')");
    // 4: required string Password (api.body="password", api.vd="regexp(^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@#$%^&*])[A-Za-z\d@#$%^&*]{8,16}$')");
    4: required string Password (api.body="password", api.vd="regexp('.{4,16}')");
}

struct MobileLoginReq {
    1: required string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    2: required string SmsCode (api.body="sms_code", api.vd="regexp('^\\d{6}$')");
}

struct MiniProgLoginReq {
    1: required string WxCode (api.body="wx_code");
}

struct PwdLoginReq {
    1: required string Username (api.body="username", api.vd="regexp('^[a-z0-9_]{1,30}$')");
//    2: required string Email (api.body="email", api.vd="regexp('^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$')");
    3: required string Password (api.body="password", api.vd="regexp('.{8,16}')");
}

struct AuthInfo {
    1: required i32 Id;
    2: required string Token;
    3: required i32 ExpiredAt;
}

struct UsernameRegisterReq {
    1: required string Username (api.body="username",  api.vd="regexp('^[a-z0-9_]{1,30}$')");
    2: required string Password (api.body="password", api.vd="regexp('.{8,16}')");
}

service auth {
    common.Empty SendSms(1: SendSmsReq req) (api.get="/api/auth/send_sms");
    AuthInfo MobileRegister(1: MobileRegisterReq req) (api.post="/api/auth/mobile_register");
    AuthInfo UsernameRegister(1:UsernameRegisterReq req) (api.post="/api/auth/username_register");
    AuthInfo MobileLogin(1: MobileLoginReq req) (api.get="/api/auth/mobile_login");
    AuthInfo MiniProgLogin(1: MiniProgLoginReq req) (api.get="/api/auth/mini_prog_login");
    AuthInfo PwdLogin(1: PwdLoginReq req) (api.get="/api/auth/pwd_login");
}

