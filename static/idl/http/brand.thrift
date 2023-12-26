namespace go brand
include "common.thrift"

struct BrandInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
    3: required string Logo (api.body="logo");
}

struct NewBrand {
    2: required string Name (api.body="name");
    3: required string Logo (api.body="logo");
}

struct GetBrandByCatReq {
    1: required i32 CategoryId (api.query="category_id");
}

service Brand {
    list<BrandInfo> GetBrandList(1: common.PageFilter req) (api.get="/api/products/brands/list");
    list<BrandInfo> GetBrandByCategory(1: GetBrandByCatReq req) (api.get="/api/products/categories/brands/list");
    BrandInfo GetBrandDetail (1:common.Req req) (api.get="/api/products/brands/detail")
    BrandInfo AddNewBrand(1: NewBrand req) (api.post="/api/products/brands/add");
    BrandInfo UpdateBrand(1: BrandInfo req) (api.put="/api/products/brands/update");
    common.Empty DelteBrand(1: common.Req req) (api.delete="/api/products/brands/delete");
}