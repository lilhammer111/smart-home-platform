namespace go banner
include "common.thrift"


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

service Banner {
    list<BannerInfo> GetAllBanners (1: common.Empty req) (api.get="/api/products/banners/list");
    BannerInfo AddNewBanner(1: NewBanner req) (api.post="/api/products/banners/add");
    BannerInfo UpdateBanner(1: BannerInfo req) (api.put="/api/products/banners/update");
    common.Empty DeleteBanner(1: common.Req req) (api.delete="/api/products/banners/delete");
}
