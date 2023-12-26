namespace go product
include "common.thrift"

struct ProductFilter {
    8: optional i16 Page (api.query="page");
    9: optional i16 Limit (api.query="limit");
    1: optional i8 State (api.query="state");
    2: optional i32 CategoryId (api.body="category_id");
    3: optional i32 BrandId (api.body="brand_id");
    4: optional list<string> Sorts (api.query="sorts");
}

struct ProductDetail {
    1: required i32 Id (api.body="id");
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

struct ProductInfo {
    1: required i32 Id (api.body="id");
    2: required i32 CategoryId (api.body="category_id");
    3: required i32 BrandId (api.body="brand_id");
    4: required string Name (api.body="name");
    6: required i8 ModelId (api.body="model_id" );
    7: required string Brief (api.body="brief");
    8: required string Picture (api.body="picture");
    9: required i8 State (api.body="state");
    10: required double Price (api.body="price");
    11: required list<string> Pictures (api.body="pictures");
}


struct RatingReq {
    2: required i32 Id (api.body="id");
    1: required double Rating (api.body="rating");
}

struct RatingResp {
    2: required i32 Id (api.body="id");
    1: required double Rating;
}

struct NewProduct {
    2: required i32 CategoryId (api.body="category_id");
    3: required i32 BrandId (api.body="brand_id");
    4: required string Name (api.body="name");
    6: required i8 ModelId (api.body="model_id");
    7: required string Brief (api.body="brief");
    8: required string Picture (api.body="picture");
    9: required i8 State (api.body="state");
    10: required double Price (api.body="price");
    1: required list<string> Showcase (api.body="showcase");
}

service Product {
    list<ProductDetail> GetProductList(1: ProductFilter req) (api.get="/api/products/list");
    ProductDetail GetProductDetail(1: common.Req req) (api.get="/api/products/detail");
    ProductInfo AddNewProduct(1: NewProduct req) (api.post="/api/products/add");
    ProductInfo UpdateProduct (1: ProductInfo req) (api.put="/api/products/update");
    RatingResp UpdateRating (1: RatingReq req) (api.put="/api/products/update_rating")
    common.Empty DeleteProduct(1: common.Req req) (api.delete="/api/products/delete");
}