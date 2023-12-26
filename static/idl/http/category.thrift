namespace go category
include "common.thrift"

struct CategoryInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
    3: required string Picture (api.body="picture");
    4: required list<string> Showcase (api.body="showcase");
}

struct AddCategoryReq {
    2: required string Name (api.body="name");
    3: required string Picture (api.body="picture");
    4: required list<string> Showcase (api.body="showcase");
}

struct CategoryBasicInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
    3: required string Picture (api.body="picture");
}

struct CategoryDetail {
    2: required i32 Id (api.body="id");
    1: required list<string> Showcase (api.body="showcase");
}



service Category {
    list<CategoryBasicInfo> GetCategoryList(1: common.PageFilter req) (api.get="/api/products/categories/list");
    CategoryDetail GetCategoryDetail(1: common.Req req) (api.get="/api/products/categories/detail");
    CategoryInfo AddNewCategory(1: AddCategoryReq req) (api.post="/api/products/categories/add");
    CategoryInfo UpdateCategory(1: CategoryInfo req) (api.put="/api/products/categories/update");
    common.Empty DeleteCategory(1: common.Req req) (api.delete="/api/products/categories/delete");
}