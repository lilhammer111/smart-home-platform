namespace go common_http

struct Resp {
    1: required bool Success;
    2: required i16 Code;
    3: required string Message;
}

struct Req {
    1: required i32 Id
}






