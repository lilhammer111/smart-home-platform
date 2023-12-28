namespace go product
include "common.thrift"
include "category.thrift"
include "brand.thrift"
include "model.thrift"

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
    11: required i32 RatingId (api.body="rating_id");
    12: required list<string> Showcase (api.body="showcase");
    13: required category.CategoryInfo Category (api.body="category");
    14: required brand.BrandInfo Brand (api.body="brand");
    15: required model.ModelInfo Model (api.body="model");
    16: required RatingInfo Rating (api.body="rating");
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
    20: required bool OnSale (api.body="on_sale");
    21: required bool IsFreeShipping (api.body="is_free_shipping");
    22: required bool IsNew (api.body="is_new");
    23: required bool IsHot (api.body="is_hot");
    24: required bool IsRecommended (api.body="is_recommended");
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
    11: required i32 RatingId (api.body="rating_id");
    12: required list<string> Pictures (api.body="pictures");
}


struct RatingReq {
    2: required i32 ProductId (api.body="product_id");
    1: required double Rating (api.body="rating");
}

struct RatingInfo {
    1: required i32 Id (api.body="id");
    2: required i32 ProductId (api.body="product_id");
    3: required i32 TotalCustomer (api.body="total_customer");
    4: required double CurRating (api.body="cur_rating");
}

service Product {
    list<ProductDetail> GetProductList(1: ProductFilter req) (api.get="/api/products/list");
    ProductDetail GetProductDetail(1: common.Req req) (api.get="/api/products/detail");
    ProductInfo AddNewProduct(1: NewProduct req) (api.post="/api/products/add");
    ProductInfo UpdateProduct (1: ProductInfo req) (api.put="/api/products/update");
    RatingInfo UpdateRating (1: RatingReq req) (api.put="/api/products/rating/update")
    common.Empty DeleteProduct(1: common.Req req) (api.delete="/api/products/delete");
}