namespace go std_resp

struct StdResp {
    1: required bool Success;
    2: required i16 Code;
    3: required string Message;
    4: optional DataUnion Data;
}

union DataUnion {
    1: CustomMap Object;
    2: list<CustomMap> Array;
}

struct Any {
    1: optional string StringValue;
    2: optional i32 IntValue;
    3: optional bool BoolValue;
    4: optional list<string> ListValue;
    5: optional map<string, Any> MapValue;
}

typedef map<string, Any> CustomMap