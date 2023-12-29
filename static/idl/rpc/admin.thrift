namespace go micro_admin

struct FoodEnumInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
}

struct AddFoodEnumReq {
    1: required i32 Id (api.body="id");

}

service EnumService {

}