namespace go feed_programs
include "std_resp.thrift"
include "std_req.thrift"

struct ProgramInfo {
    1: required i8 Frequency (api.body="frequency");
    2: required i8 FeedingAmount (api.body="feeding_amount");
    3: required i16 FeedingTime (api.body="feeding_time");
    4: optional i32 Id (api.body="id");
    5: required i32 DeviceId (api.body="device_id");
    6: required string FoodType (api.body="food_type");
}

struct FeedNowReq {
    1: required i32 DeviceId (api.body="device_id");
    2: required i8 FeedingAmount (api.body="feeding_amount");
    3: required string FoodType (api.body="food_type");
}

service feed_programs {
    std_resp.StdResp GetSelfProgramList(1: req.IdReq req) (api.get="/api/programs/list");
    std_resp.StdResp GetProgramDetail(1: req.IdReq req) (api.get="/api/programs/detail");
    std_resp.StdResp UpdateProgramInfo(1: ProgramInfo req) (api.put="/api/programs/update");
    std_resp.StdResp CreateProgram(1: ProgramInfo req) (api.post="/api/programs/create");
    std_resp.StdResp FeedNow(1: FeedNowReq req) (api.post="/api/programs/feed_now");
    std_resp.StdResp DeleteProgram(1: req.IdReq req) (api.delete="/api/programs/delete");
}