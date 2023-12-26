namespace go common


struct Req {
    1: required i32 Id
}

struct Empty {
}

struct PageFilter {
    1: optional i16 Page (api.query="page");
    2: optional i16 Limit (api.query="limit");
}




