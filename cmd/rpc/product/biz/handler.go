package biz

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/banner_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/brand_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/category_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/model_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/product_service"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

// CombineServiceImpl implements the last service interface defined in the IDL.
type CombineServiceImpl struct{}

// GetProductList implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) GetProductList(ctx context.Context, req *product.ProductFilter) (resp []*product.BasicProdInfo, err error) {
	resp, err = product_service.NewGetProductListService(ctx).Run(req)

	return resp, err
}

// GetProductDetail implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) GetProductDetail(ctx context.Context, req *common.Req) (resp *product.ProductInfo, err error) {
	resp, err = product_service.NewGetProductDetailService(ctx).Run(req)

	return resp, err
}

// AddNewProduct implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) AddNewProduct(ctx context.Context, req *product.NewProduct_) (resp *product.ProductInfo, err error) {
	resp, err = product_service.NewAddNewProductService(ctx).Run(req)

	return resp, err
}

// UpdateProductBasicInfo implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) UpdateProductBasicInfo(ctx context.Context, req *product.BasicProdInfo) (resp *product.BasicProdInfo, err error) {
	resp, err = product_service.NewUpdateProductBasicInfoService(ctx).Run(req)

	return resp, err
}

// UpdateProdShowcase implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) UpdateProdShowcase(ctx context.Context, req *product.ProductShowcase) (resp *product.ProductShowcase, err error) {
	resp, err = product_service.NewUpdateProdShowcaseService(ctx).Run(req)

	return resp, err
}

// UpdateRating implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) UpdateRating(ctx context.Context, req *product.RatingReq) (resp *product.RatingResp, err error) {
	resp, err = product_service.NewUpdateRatingService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductServiceImpl interface.
func (s *CombineServiceImpl) DeleteProduct(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = product_service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// GetCategoryList implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) GetCategoryList(ctx context.Context, req *product.PageFilter) (resp []*product.CategoryBasicInfo, err error) {
	resp, err = category_service.NewGetCategoryListService(ctx).Run(req)

	return resp, err
}

// GetCategoryDetail implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) GetCategoryDetail(ctx context.Context, req *common.Req) (resp *product.CategoryDetail, err error) {
	resp, err = category_service.NewGetCategoryDetailService(ctx).Run(req)

	return resp, err
}

// AddNewCategory implements the CategoryServiceImpl interface.
func (s *CombineServiceImpl) AddNewCategory(ctx context.Context, req *product.AddCategoryReq) (resp *product.CategoryInfo, err error) {
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

// GetAllModels implements the ModelServiceImpl interface.
func (s *CombineServiceImpl) GetAllModels(ctx context.Context) (resp []*product.ModelInfo, err error) {
	resp, err = model_service.NewGetAllModelsService(ctx).Run()

	return resp, err
}

// AddNewModel implements the ModelServiceImpl interface.
func (s *CombineServiceImpl) AddNewModel(ctx context.Context, req *product.NewModel_) (resp *product.ModelInfo, err error) {
	resp, err = model_service.NewAddNewModelService(ctx).Run(req)

	return resp, err
}

// DeleteModel implements the ModelServiceImpl interface.
func (s *CombineServiceImpl) DeleteModel(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = model_service.NewDeleteModelService(ctx).Run(req)

	return resp, err
}

// GetBrandList implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) GetBrandList(ctx context.Context, req *product.PageFilter) (resp []*product.BrandInfo, err error) {
	resp, err = brand_service.NewGetBrandListService(ctx).Run(req)

	return resp, err
}

// GetBrandByCategory implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) GetBrandByCategory(ctx context.Context, req *product.GetBrandByCatReq) (resp []*product.BrandInfo, err error) {
	resp, err = brand_service.NewGetBrandByCategoryService(ctx).Run(req)

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

// DelteBrand implements the BrandServiceImpl interface.
func (s *CombineServiceImpl) DelteBrand(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = brand_service.NewDelteBrandService(ctx).Run(req)

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

// DelteBanner implements the BannerServiceImpl interface.
func (s *CombineServiceImpl) DelteBanner(ctx context.Context, req *common.Req) (resp *common.Empty, err error) {
	resp, err = banner_service.NewDelteBannerService(ctx).Run(req)

	return resp, err
}
