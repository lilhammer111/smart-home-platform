namespace go model
include "common.thrift"


struct ModelInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
}

struct NewModel {
    1: required string Name (api.body="name");
}

service Model {
    list<ModelInfo> GetAllModels(1: common.Empty req) (api.get="/api/products/models/list");
    ModelInfo GetModelDetail(1: common.Req req) (api.get="/api/products/models/detail");
    ModelInfo AddNewModel(1: NewModel req) (api.post="/api/products/models/add");
    common.Empty DeleteModel(1: common.Req req) (api.delete="/api/products/models/delete");
}