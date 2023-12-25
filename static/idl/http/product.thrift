namespace go product
include "common.thrift"

struct ProductFilter {
    1: optional i8 State (api.query="state");
    2: required i32 CategoryId (api.body="category_id");
    3: required i32 BrandId (api.body="brand_id");
    4: optional list<string> Sorts (api.query="sorts");
    6: optional string StartDate (api.query="start_date");
    7: optional string EndDate (api.query="end_date");
}

struct ProductInfo {
    1: optional i32 Id (api.body="id");
    2: required i32 CategoryId (api.body="category_id");
    3: required i32 BrandId (api.body="brand_id");
    4: required string Name (api.body="name");
    6: required i8 ModelId (api.body="model_id" );
    7: required string Brief (api.body="brief");
    8: required string Picture (api.body="picture");
    9: required i8 State (api.body="state");
    10: required double Price (api.body="price");
    11: required double Rating (api.body="rating");
    12: required list<string> Showcase (api.body="showcase");
}

struct CategoryInfo {

}


service product {
    list<ProductInfo > GetProductList(1: ProductFilter req) (api.get="/api/products/list");
    ProductInfo  GetProductDetail(1: common.Req req) (api.get="/api/products/detail");
    ProductInfo  UpdateProductInfo (1: ProductInfo  req) (api.put="/api/products/update");
    ProductInfo  CreateProduct(1: ProductInfo  req) (api.post="/api/products/create");
    common.Empty DeleteProduct(1: common.Req req) (api.delete="/api/products/delete");
}

service category {
    list<ProductInfo > GetProductList(1: ProductFilter req) (api.get="/api/products/list");
    ProductInfo  GetProductDetail(1: common.Req req) (api.get="/api/products/detail");
    ProductInfo  UpdateProductInfo (1: ProductInfo  req) (api.put="/api/products/update");
    ProductInfo  CreateProduct(1: ProductInfo  req) (api.post="/api/products/create");
    common.Empty DeleteProduct(1: common.Req req) (api.delete="/api/products/delete");
}