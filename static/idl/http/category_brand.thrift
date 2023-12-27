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
    list<CategoryBrandInfo> BatchAddCategoryBrand(1: NewCategoryBrand req) (api.post="/api/products/category_brand/batch_add");
    common.Empty BatchReduceCategoryBrand(1: NewCategoryBrand req) (api.delete="/api/products/category_brand/batch_reduce");
    common.Empty DeleteBrandByCategory(1: common.Req req) (api.delete="/api/products/category_brand/delete_brand");
    common.Empty DeleteCategoryByBrand(1: common.Req req) (api.delete="/api/products/category_brand/delete_category");
}