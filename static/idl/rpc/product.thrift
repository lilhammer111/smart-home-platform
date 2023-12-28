namespace go product
include "../http/common.thrift"
include "../http/category.thrift"
include "../http/brand.thrift"
include "../http/model.thrift"

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
    6: required string Price (api.body="price");
    7: required string Rating (api.body="rating");
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

struct ProductDetailResp {
    1: required i32 Id (api.body="id");
    2: optional string Name (api.body="name");
    3: optional string Brief (api.body="brief");
    4: optional string Picture (api.body="picture");
    6: optional string Price (api.body="price");
    5: optional ProductState State (api.body="state");

    10: optional i32 CategoryId (api.body="category_id");
    12: optional i32 BrandId (api.body="brand_id");
    13: optional list<string> ModelList (api.body="model_list");
    14: optional list<string> Showcase (api.body="showcase");

    7: required string Rating (api.body="rating");
}

struct ProductInfo {
    1: required i32 Id (api.body="id");
    2: optional string Name (api.body="name");
    3: optional string Brief (api.body="brief");
    4: optional string Picture (api.body="picture");
    6: optional string Price (api.body="price");
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

service ProductService {
    list<ProductBasicInfo> GetProductList(1: ProductFilter req) (api.get="/api/products/list");
    ProductDetailResp GetProductDetail(1: common.Req req) (api.get="/api/products/detail");
    ProductInfo AddNewProduct(1: AddProductReq req) (api.post="/api/products/add");
    ProductInfo UpdateProduct (1: ProductInfo req) (api.put="/api/products/update");
    RatingResp RateProduct (1: RatingReq req) (api.put="/api/products/rating/update")
    common.Empty DeleteProduct(1: common.Req req) (api.delete="/api/products/delete");
}

struct CategoryInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
    3: required string Picture (api.body="picture");
    4: required list<string> Showcase (api.body="showcase");
    5: required string Brief (api.body="brief")
}

struct NewCategory {
    2: required string Name (api.body="name");
    3: required string Picture (api.body="picture");
    4: required list<string> Showcase (api.body="showcase");
    5: required string Brief (api.body="brief")
}

service CategoryService {
    list<CategoryInfo> GetCategoryList(1: common.PageFilter req) (api.get="/api/products/categories/list");
    CategoryInfo GetCategoryDetail(1: common.Req req) (api.get="/api/products/categories/detail");
    CategoryInfo AddNewCategory(1: NewCategory req) (api.post="/api/products/categories/add");
    CategoryInfo UpdateCategory(1: CategoryInfo req) (api.put="/api/products/categories/update");
    common.Empty DeleteCategory(1: common.Req req) (api.delete="/api/products/categories/delete");
}

struct ModelInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
}

struct NewModel {
    1: required string Name (api.body="name");
}

service ModelService {
    list<ModelInfo> GetAllModels() (api.get="/api/products/models/list");
    ModelInfo GetModelDetail(1: common.Req req) (api.get="/api/products/models/detail");
    ModelInfo AddNewModel(1: NewModel req) (api.post="/api/products/models/add");
    common.Empty DeleteModel(1: common.Req req) (api.delete="/api/products/models/delete");
}

struct BrandInfo {
    1: required i32 Id (api.body="id");
    2: required string Name (api.body="name");
    3: required string Logo (api.body="logo");
}

struct NewBrand {
    2: required string Name (api.body="name");
    3: required string Logo (api.body="logo");
}

struct BrandByCatReq {
    1: required i32 CategoryId (api.query="category_id");
}

service BrandService {
    list<BrandInfo> GetBrandList(1: common.PageFilter req) (api.get="/api/products/brands/list");
    list<BrandInfo> GetRelatedBrandsByCategoryId(1: BrandByCatReq req) (api.get="/api/products/brands/related");
    BrandInfo GetBrandDetail (1:common.Req req) (api.get="/api/products/brands/detail")
    BrandInfo AddNewBrand(1: NewBrand req) (api.post="/api/products/brands/add");
    BrandInfo UpdateBrand(1: BrandInfo req) (api.put="/api/products/brands/update");
    common.Empty DeleteBrand(1: common.Req req) (api.delete="/api/products/brands/delete");
}

struct BannerInfo {
    1: required i8 Id (api.body="id");
    2: required string Picture(api.body="picture");
    3: required string ProductLink(api.body="product_link");
    4: required i8  Index (api.body="index");
}

struct NewBanner {
    2: required string Picture(api.body="picture");
    3: required string ProductLink(api.body="product_link");
    4: required i8  Index (api.body="index");
}

service BannerService {
    list<BannerInfo> GetAllBanners () (api.get="/api/products/banners/list");
    BannerInfo AddNewBanner(1: NewBanner req) (api.post="/api/products/banners/add");
    BannerInfo UpdateBanner(1: BannerInfo req) (api.put="/api/products/banners/update");
    common.Empty DeleteBanner(1: common.Req req) (api.delete="/api/products/banners/delete");
}

struct CategoryBrandInfo {
    1: required i32 Id;
    2: required i32 CategoryId;
    3: required i32 BrandId;
}

struct NewCategoryBrand {
    1: required i32 BrandId;
    2: required list<i32> CategoryId;
}

service CategoryBrandService {
    list<CategoryBrandInfo> BatchAddCategoryBrand(1: NewCategoryBrand req) (api.post="/api/products/category_brand/batch_add");
    common.Empty BatchReduceCategoryBrand(1: NewCategoryBrand req) (api.delete="/api/products/category_brand/batch_reduce");
    common.Empty DeleteBrandByCategory(1: common.Req req) (api.delete="/api/products/category_brand/delete_brand");
    common.Empty DeleteCategoryByBrand(1: common.Req req) (api.delete="/api/products/category_brand/delete_category");
}