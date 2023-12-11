namespace go auth
include "std_resp.thrift"

struct SendSmsReq {
    1: required string Mobile (api.query="mobile", api.vd="regexp('^1[3-9]\\d{9}$)");
}

struct MobileRegisterReq {
    1: required string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    2: required string SmsCode (api.body="sms_code", api.vd="regexp('^\\d{6}$')");
    3: required string Username (api.body="username", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')");
    // 4: required string Password (api.body="password", api.vd="regexp(^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@#$%^&*])[A-Za-z\d@#$%^&*]{8,16}$')");
    4: required string Password (api.body="password", api.vd="regexp('.{4,16}')");
}

struct MobileLoginReq {
    1: required string Mobile (api.body="mobile", api.vd="regexp('^1[3-9]\\d{9}$')");
    2: required string SmsCode (api.body="sms_code", api.vd="regexp('^\\d{6}$')");
}

struct MiniProgLoginReq {
    1: required string WxCode;
}

struct PwdLoginReq {
    1: required string Username (api.body="username", api.vd="regexp('^[\p{L}\p{N}\p{Lo}_]{1,15}$')");
    2: required string Password (api.body="password", api.vd="regexp('.{4,16}')");
}

service auth {
    std_resp.StdResp SendSms(1: SendSmsReq req) (api.get="/api/auth/send_sms");
    std_resp.StdResp MobileRegister(1: MobileRegisterReq req) (api.post="/api/auth/mobile_register");
    std_resp.StdResp MobileLogin(1: MobileLoginReq req) (api.get="/api/auth/mobile_login");
    std_resp.StdResp MiniProgLogin(1: MiniProgLoginReq req) (api.get="/api/auth/mini_prog_login");
    std_resp.StdResp PwdLogin(1: PwdLoginReq req) (api.get="/api/auth/pwd_login");
}

