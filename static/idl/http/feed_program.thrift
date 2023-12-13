namespace go feed_program
include "../common_http.thrift"


struct ProgramInfo {
    1: optional i32 Id (api.body="id");
    2: required i8 Frequency (api.body="frequency");
    3: required i16 Amount (api.body="amount");
    4: required string FeedTime (api.body="feed_time");
    5: required i32 DeviceId (api.body="device_id");
    6: required string Food (api.body="food");
}

struct FeedNowReq {
    1: required i32 DeviceId (api.body="device_id");
    2: required i8 Amount (api.body="amount");
    3: required string Food (api.body="food");
}

// response
struct ProgramListResp {
    1: required common_http.Resp Meta;
    2: required list<i32> Data;
}

struct ProgramInfoResp {
    1: required common_http.Resp Meta;
    2: required ProgramInfo Data;
}


service feed_program {
    ProgramListResp GetSelfProgramList(1: common_http.Req req) (api.get="/api/programs/list");
    ProgramInfoResp GetProgramDetail(1: common_http.Req req) (api.get="/api/programs/detail");
    ProgramInfoResp UpdateProgramInfo(1: ProgramInfo req) (api.put="/api/programs/update");
    ProgramInfoResp CreateProgram(1: ProgramInfo req) (api.post="/api/programs/create");
    common_http.Resp FeedNow(1: FeedNowReq req) (api.post="/api/programs/feed_now");
    common_http.Resp DeleteProgram(1: common_http.Req req) (api.delete="/api/programs/delete");
}