namespace go product
include "common.thrift"
include "category.thrift"
include "brand.thrift"
include "model.thrift"

struct ProductFilter {
    30: optional i16 Page (api.query="page");
    31: optional i16 Limit (api.query="limit");

    1: optional string Search (api.query="search")
    2: optional i32 CategoryId (api.query="category_id");
    3: optional list<i32> BrandIdList (api.query="brand_id_list");
    4: optional bool IsPriceAsc (api.query="is_price_asc");
    5: optional bool IsRatingAsc (api.query="is_rating_asc");

    20: optional bool OnSale (api.query="on_sale");
    21: optional bool IsFreeShipping (api.query="is_free_shipping");
    22: optional bool IsNew (api.query="is_new");
    23: optional bool IsHot (api.query="is_hot");
    24: optional bool IsRecommended (api.query="is_recommended");
}


struct ProductState {
    20: required bool OnSale (api.body="on_sale");
    21: required bool IsFreeShipping (api.body="is_free_shipping");
    22: required bool IsNew (api.body="is_new");
    23: required bool IsHot (api.body="is_hot");
    24: required bool IsRecommended (api.body="is_recommended");
}

struct OptionalState {
    20: optional bool OnSale (api.body="on_sale");
    21: optional bool IsFreeShipping (api.body="is_free_shipping");
    22: optional bool IsNew (api.body="is_new");
    23: optional bool IsHot (api.body="is_hot");
    24: optional bool IsRecommended (api.body="is_recommended");
}

struct ProductBasicInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
    3: required string Brief (api.body="brief");
    4: required string Picture (api.body="picture");
    6: required double Price (api.body="price");
    7: required double Rating (api.body="rating");
}

struct ProductDetailResp {
    1: required i32 Id (api.body="id");
    2: optional string Name (api.body="name");
    3: optional string Brief (api.body="brief");
    4: optional string Picture (api.body="picture");
    6: optional double Price (api.body="price");
    5: optional ProductState State (api.body="state");

    10: optional i32 CategoryId (api.body="category_id");
    12: optional i32 BrandId (api.body="brand_id");
    13: optional list<string> ModelList (api.body="model_list");
    14: optional list<string> Showcase (api.body="showcase");

    7: required double Rating (api.body="rating");
}

struct AddProductReq {
    2: required string Name (api.body="name");
    3: required string Brief (api.body="brief");
    4: required string Picture (api.body="picture");
    6: required double Price (api.body="price");
    5: required ProductState State (api.body="state");

    10: required i32 CategoryId (api.body="category_id");
    12: required i32 BrandId (api.body="brand_id");
    13: required list<string> ModelList (api.body="model_list");
    14: required list<string> Showcase (api.body="showcase");
}

struct ProductInfo {
    1: required i32 Id (api.body="id");
    2: optional string Name (api.body="name");
    3: optional string Brief (api.body="brief");
    4: optional string Picture (api.body="picture");
    6: optional double Price (api.body="price");
    5: optional OptionalState State (api.body="state");

    10: optional i32 CategoryId (api.body="category_id");
    12: optional i32 BrandId (api.body="brand_id");
    13: optional list<string> ModelList (api.body="model_list");
    14: optional list<string> Showcase (api.body="showcase");
}

struct RatingReq {
    2: required i32 ProductId (api.body="product_id");
    1: required double Rating (api.body="rating");
}

struct RatingResp {
    4: required string Rating (api.body="rating");
}

service Product {
    list<ProductBasicInfo> GetProductList(1: ProductFilter req) (api.get="/api/products/list");
    ProductDetailResp GetProductDetail(1: common.Req req) (api.get="/api/products/detail");
    ProductInfo AddNewProduct(1: AddProductReq req) (api.post="/api/products/add");
    ProductInfo UpdateProduct (1: ProductInfo req) (api.put="/api/products/update");
    RatingResp RateProduct (1: RatingReq req) (api.put="/api/products/rating/update")
    common.Empty DeleteProduct(1: common.Req req) (api.delete="/api/products/delete");
}