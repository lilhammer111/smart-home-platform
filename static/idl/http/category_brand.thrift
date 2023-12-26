namespace go category_brand
include "common.thrift"

struct CategoryBrandInfo {
    1: required i32 Id (api.body="id");
    2: required i32 CategoryId (api.body="category_id");
    3: required i32 BrandId (api.body="brand_id");
}

struct NewCategoryBrand {
    1: required i32 BrandId (api.body="brand_id");
    2: required list<i32> CategoryId (api.body="category_id");
}


service CategoryBrand {
    list<CategoryBrandInfo> GetCategoryBrandList(1: common.Req req) (api.get="/api/products/category_brand/list");
    list<CategoryBrandInfo> BatchAddCategoryBrand(1: NewCategoryBrand req) (api.post="/api/products/category_brand/batch_add");
    CategoryBrandInfo UpdateCategoryBrand(1: CategoryBrandInfo req) (api.put="/api/products/category_brand/update");
    common.Empty DeleteCategoryByBrand(1: common.Req req) (api.delete="/api/products/category_brand/delete_category");
    common.Empty DeleteBrandByCategory(1: common.Req req) (api.delte="/api/products/category_brand/delete_brand")
}