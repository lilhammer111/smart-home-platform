namespace go feed_program
include "standard.thrift"


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

service feed_program {
    standard.Resp GetSelfProgramList(1: i32 req) (api.get="/api/programs/list");
    standard.Resp GetProgramDetail(1: i32 req) (api.get="/api/programs/detail");
    standard.Resp UpdateProgramInfo(1: ProgramInfo req) (api.put="/api/programs/update");
    standard.Resp CreateProgram(1: ProgramInfo req) (api.post="/api/programs/create");
    standard.Resp FeedNow(1: FeedNowReq req) (api.post="/api/programs/feed_now");
    standard.Resp DeleteProgram(1: i32 req) (api.delete="/api/programs/delete");
}