namespace go product
include "../http/common.thrift"
include "../http/category.thrift"
include "../http/brand.thrift"
include "../http/model.thrift"

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
    13: required category.CategoryInfo Category (api.body="category");
    14: required brand.BrandInfo Brand (api.body="brand");
    15: required model.ModelInfo Model (api.body="model");
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

service ProductService {
    list<ProductDetail> GetProductList(1: ProductFilter req) (api.get="/api/products/list");
    ProductDetail GetProductDetail(1: common.Req req) (api.get="/api/products/detail");
    ProductInfo AddNewProduct(1: NewProduct req) (api.post="/api/products/add");
    ProductInfo UpdateProduct (1: ProductInfo req) (api.put="/api/products/update");
    RatingResp UpdateRating (1: RatingReq req) (api.put="/api/products/update_rating")
    common.Empty DeleteProduct(1: common.Req req) (api.delete="/api/products/delete");
}


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

struct PageFilter {
    1: optional i16 Page (api.query="page");
    2: optional i16 Limit (api.query="limit");
}

service CategoryService {
    list<CategoryBasicInfo> GetCategoryList(1: PageFilter req) (api.get="/api/products/categories/list");
    CategoryDetail GetCategoryDetail(1: common.Req req) (api.get="/api/products/categories/detail");
    CategoryInfo AddNewCategory(1: AddCategoryReq req) (api.post="/api/products/categories/add");
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
    list<BrandInfo> GetBrandList(1: PageFilter req) (api.get="/api/products/brands/list");
    list<BrandInfo> GetBrandListByCategory(1: BrandByCatReq req) (api.get="/api/products/categories/brands/list");
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
    list<CategoryBrandInfo> GetCategoryBrandList(1: common.Req req) (api.get="/api/products/category_brand/list");
    list<CategoryBrandInfo> BatchAddCategoryBrand(1: NewCategoryBrand req) (api.post="/api/products/category_brand/batch_add");
    CategoryBrandInfo UpdateCategoryBrand(1: CategoryBrandInfo req) (api.put="/api/products/category_brand/update");
    common.Empty DeleteCategoryByBrand(1: common.Req req) (api.delete="/api/products/category_brand/delete_category");
    common.Empty DeleteBrandByCategory(1: common.Req req) (api.delte="/api/products/category_brand/delete_brand")
}