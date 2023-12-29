namespace go micro_feeder
include "../http/common.thrift"




struct ProgramInfo {
    1: required i32 Id (api.body="id");
    7: required i32 FeederId (api.body="feeder_id")
    2: required list<i32> Frequency (api.body="frequency");
    3: required i16 FeedTime (api.body="feed_time");
    4: required i16 Amount (api.body="amount", api.vd="$ > 0 && $ <= 100");
    5: required i32 DeviceId (api.vd="$ > 0 && $ <= 2147483647");
    6: required i8 Food (api.vd="$ >= 0 && $ <= 127");
}

struct UpdateProgramInfoReq {
    1: required i32 Id (api.body="id", api.vd="$>0;msg:Invalid Id Parameter.");
    7: optional i32 FeederId (api.body="feeder_id", api.vd="$>0;msg:Invalid Feeder Id");
    2: optional list<i32> Frequency (api.body="frequency", api.vd="range($, in(#v, Mon,Tus,Wed,Thur,Fri,Sat,Sun,Everyday,Weekend,Weekday));msg:Invalid Frequency Parameter.");
    3: optional i16 FeedTime (api.body="feed_time", api.vd="$ > 0 && ($ / 100) < 24 && ($ % 100) < 60;msg:Invalid Feed Time");
    4: optional i16 Amount (api.body="amount", api.vd="$ > 0 && $ <= 100 && ($ % 10 == 0);msg:Invalid Amount");
    6: optional i8 Food (api.body="food", api.vd="$ >= 0 && $ <= 127;msg:Invalid Food Parameter");
}

struct AddProgramReq {
    7: required i32 FeederId (api.body="feeder_id")
    2: required list<i32> Frequency (api.body="frequency", api.vd="range($, in(#v, Mon,Tus,Wed,Thur,Fri,Sat,Sun,Everyday,Weekend,Weekday));msg:Invalid Frequency Parameter.");
    3: required i16 FeedTime (api.body="feed_time", api.vd="($ / 100) < 24 && ($ / 100) > 0 && ($ % 100) < 60 ($ % 100) > 0");
    4: required i16 Amount (api.body="amount", api.vd="$ > 0 && $ <= 100");
    5: required i32 DeviceId (api.vd="$ > 0 && $ <= 2147483647");
    6: required i8 Food (api.vd="$ >= 0 && $ <= 127");
}

service ProgramService {
    list<ProgramInfo> GetProgramList(1: common.Empty req) (api.get="/api/feeder/program/list");
    ProgramInfo AddProgram(1: AddProgramReq req) (api.post="/api/feeder/program/add");
    ProgramInfo UpdateProgram(1: UpdateProgramInfoReq req) (api.put="/api/feeder/program/update");
    common.Empty DeleteProgram(1: common.Req req) (api.delete="/api/feeder/program/delete");
}

struct RecordFilter {
    1: optional i32 FeederId (api.query="feeder_id");
    2: optional bool IsSuccessful (api.query="is_successful");
    3: optional i8 Food (api.query="food");
    4: optional i16 GtAmount (api.query="gt_amount");
    5: optional i16 LtAmount (api.query="lt_amount");
    6: optional i32 FedAtStart (api.query="fed_at_start");
    7: optional i32 FedAtEnd (api.query="fed_at_end");
    8: optional string Search (api.query="search");
}

struct RecordInfo {
    7: required i32 Id;
    1: required i32 FeederId;
    2: required bool IsSuccessful;
    3: required i8 Food;
    4: required i16 Amount;
    5: required i32 FedAt;
    6: required string Remark;
}

struct AddRecordReq {
    1: required i32 FeederId (api.body="feeder_id", api.vd="$>0;msg:Invalid Feeder Id");
    2: required bool IsSuccessful (api.body="is_successful", api.vd="typeof($) == 'boolean';msg:Invalid IsSuccessful Parameter");
    3: required i8 Food (api.body="food", api.vd="$ >= 0 && $ <= 127;msg:Invalid Food Parameter");
    4: required i16 Amount (api.body="amount", api.vd="$ > 0 && $ <= 100;msg:Invalid Amount");
    5: required i32 FedAt (api.body="fed_at", api.vd="$>0;msg:Invalid FedAt Parameter");
    6: required string Remark (api.body="remark", api.vd="len($) <= 255;msg:Remark Too Long");
}

service RecordService {
    list<RecordInfo> GetRecordList(1: RecordFilter req) (api.get="/api/feeder/record/list");
    RecordInfo PostRecord(1: AddRecordReq req) (api.post="/api/feeder/record/post");
    common.Empty DeleteRecord(1: common.Req req) (api.delete="/api/feeder/record/delete");
}

struct FeederInfo {
    1: required i32 DeviceId;
    6: required i16 FoodSurplus;
    7: required i8 BatteryCharge;
}

struct UpdateFeederReq {
    1: required i32 Id;
    2: required i32 DeviceId;
    6: optional i16 FoodSurplus;
    7: optional i8 BatteryCharge;
}

service FeederService {
    FeederInfo GetFeederDetail(1: common.Req req) (api.get="/api/feeder/detail");
    FeederInfo InitFeeder(1: common.Req req) (api.post="/api/feeder/init");
    FeederInfo UpdateFeeder(1: UpdateFeederReq req) (api.put="/api/feeder/update");
    common.Empty DeleteFeeder(1: common.Req req) (api.delete="/api/feeder/delete");
}


