package biz

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/banner_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/brand_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/category_brand_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/category_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/product_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}

// GetProductList implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) GetProductList(ctx context.Context, req *product.ProductFilter) (resp []*product.ProductBasicInfo, err error) {
	resp, err = product_service.NewGetProductListService(ctx).Run(req)

	return resp, err
}

// GetProductDetail implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) GetProductDetail(ctx context.Context, req *common.Req) (resp *product.ProductDetailResp, err error) {
	resp, err = product_service.NewGetProductDetailService(ctx).Run(req)

	return resp, err
}

// AddNewProduct implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) AddNewProduct(ctx context.Context, req *product.AddProductReq) (resp *product.ProductInfo, err error) {
	resp, err = product_service.NewAddNewProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) UpdateProduct(ctx context.Context, req *product.ProductInfo) (resp *product.ProductInfo, err error) {
	resp, err = product_service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) DeleteProduct(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = product_service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// GetCategoryList implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) GetCategoryList(ctx context.Context, req *common.PageFilter) (resp []*product.CategoryInfo, err error) {
	resp, err = category_service.NewGetCategoryListService(ctx).Run(req)

	return resp, err
}

// GetCategoryDetail implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) GetCategoryDetail(ctx context.Context, req *common.Req) (resp *product.CategoryInfo, err error) {
	resp, err = category_service.NewGetCategoryDetailService(ctx).Run(req)

	return resp, err
}

// AddNewCategory implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) AddNewCategory(ctx context.Context, req *product.NewCategory_) (resp *product.CategoryInfo, err error) {
	resp, err = category_service.NewAddNewCategoryService(ctx).Run(req)

	return resp, err
}

// UpdateCategory implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) UpdateCategory(ctx context.Context, req *product.CategoryInfo) (resp *product.CategoryInfo, err error) {
	resp, err = category_service.NewUpdateCategoryService(ctx).Run(req)

	return resp, err
}

// DeleteCategory implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) DeleteCategory(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = category_service.NewDeleteCategoryService(ctx).Run(req)

	return resp, err
}

// GetBrandList implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) GetBrandList(ctx context.Context, req *common.PageFilter) (resp []*product.BrandListResp, err error) {
	resp, err = brand_service.NewGetBrandListService(ctx).Run(req)

	return resp, err
}

// GetBrandDetail implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) GetBrandDetail(ctx context.Context, req *common.Req) (resp *product.BrandInfo, err error) {
	resp, err = brand_service.NewGetBrandDetailService(ctx).Run(req)

	return resp, err
}

// AddNewBrand implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) AddNewBrand(ctx context.Context, req *product.NewBrand_) (resp *product.BrandInfo, err error) {
	resp, err = brand_service.NewAddNewBrandService(ctx).Run(req)

	return resp, err
}

// UpdateBrand implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) UpdateBrand(ctx context.Context, req *product.BrandInfo) (resp *product.BrandInfo, err error) {
	resp, err = brand_service.NewUpdateBrandService(ctx).Run(req)

	return resp, err
}

// DeleteBrand implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) DeleteBrand(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = brand_service.NewDeleteBrandService(ctx).Run(req)

	return resp, err
}

// GetAllBanners implements the BannerServiceImpl interface.
func (s *CombineServiceImpl) GetAllBanners(ctx context.Context) (resp []*product.BannerInfo, err error) {
	resp, err = banner_service.NewGetAllBannersService(ctx).Run()

	return resp, err
}

// AddNewBanner implements the BannerServiceImpl interface.
func (s *CombineServiceImpl) AddNewBanner(ctx context.Context, req *product.NewBanner_) (resp *product.BannerInfo, err error) {
	resp, err = banner_service.NewAddNewBannerService(ctx).Run(req)

	return resp, err
}

// UpdateBanner implements the BannerServiceImpl interface.
func (s *CombineServiceImpl) UpdateBanner(ctx context.Context, req *product.BannerInfo) (resp *product.BannerInfo, err error) {
	resp, err = banner_service.NewUpdateBannerService(ctx).Run(req)

	return resp, err
}

// DeleteBanner implements the BannerServiceImpl interface.
func (s *CombineServiceImpl) DeleteBanner(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = banner_service.NewDeleteBannerService(ctx).Run(req)

	return resp, err
}

// BatchAddCategoryBrand implements the CategoryBrandServiceImpl interface.
func (s *CombineServiceImpl) BatchAddCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_) (resp []*product.CategoryBrandInfo, err error) {
	resp, err = category_brand_service.NewBatchAddCategoryBrandService(ctx).Run(req)

	return resp, err
}

// DeleteCategoryByBrand implements the CategoryBrandServiceImpl interface.
func (s *CombineServiceImpl) DeleteCategoryByBrand(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = category_brand_service.NewDeleteCategoryByBrandService(ctx).Run(req)

	return resp, err
}

// DeleteBrandByCategory implements the CategoryBrandServiceImpl interface.
func (s *CombineServiceImpl) DeleteBrandByCategory(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = category_brand_service.NewDeleteBrandByCategoryService(ctx).Run(req)

	return resp, err
}

// BatchReduceCategoryBrand implements the CategoryBrandServiceImpl interface.
func (s *CombineServiceImpl) BatchReduceCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_) (resp *common.Empty, err error) {
	resp, err = category_brand_service.NewBatchReduceCategoryBrandService(ctx).Run(req)

	return resp, err
}

// GetRelatedBrandsByCategoryId implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) GetRelatedBrandsByCategoryId(ctx context.Context, req *product.BrandByCatReq) (resp []*product.BrandInfo, err error) {
	resp, err = brand_service.NewGetRelatedBrandsByCategoryIdService(ctx).Run(req)

	return resp, err
}

// RateProduct implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) RateProduct(ctx context.Context, req *product.RatingReq) (resp *product.RatingResp, err error) {
	resp, err = product_service.NewRateProductService(ctx).Run(req)

	return resp, err
}

// UpdateCategoryBrand implements the CategoryBrandServiceImpl interface.
func (s *CombineServiceImpl) UpdateCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_) (resp *common.Empty, err error) {
	resp, err = service.NewUpdateCategoryBrandService(ctx).Run(req)

	return resp, err
}
