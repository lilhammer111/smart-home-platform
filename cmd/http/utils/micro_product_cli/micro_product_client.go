package micro_product_cli

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product/combineservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() combineservice.Client
	Service() string
	GetProductList(ctx context.Context, req *product.ProductFilter, callOptions ...callopt.Option) (resp []*product.ProductDetail, err error)

	GetProductDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ProductDetail, err error)

	AddNewProduct(ctx context.Context, req *product.NewProduct_, callOptions ...callopt.Option) (resp *product.ProductInfo, err error)

	UpdateProduct(ctx context.Context, req *product.ProductInfo, callOptions ...callopt.Option) (resp *product.ProductInfo, err error)

	UpdateRating(ctx context.Context, req *product.RatingReq, callOptions ...callopt.Option) (resp *product.RatingInfo, err error)

	DeleteProduct(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error)

	GetCategoryList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (resp []*product.CategoryInfo, err error)

	GetCategoryDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error)

	AddNewCategory(ctx context.Context, req *product.NewCategory_, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error)

	UpdateCategory(ctx context.Context, req *product.CategoryInfo, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error)

	DeleteCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error)

	GetAllModels(ctx context.Context, callOptions ...callopt.Option) (resp []*product.ModelInfo, err error)

	GetModelDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ModelInfo, err error)

	AddNewModel(ctx context.Context, req *product.NewModel_, callOptions ...callopt.Option) (resp *product.ModelInfo, err error)

	DeleteModel(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error)

	GetBrandList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error)

	GetRelatedBrandsByCategoryId(ctx context.Context, req *product.BrandByCatReq, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error)

	GetBrandDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.BrandInfo, err error)

	AddNewBrand(ctx context.Context, req *product.NewBrand_, callOptions ...callopt.Option) (resp *product.BrandInfo, err error)

	UpdateBrand(ctx context.Context, req *product.BrandInfo, callOptions ...callopt.Option) (resp *product.BrandInfo, err error)

	DeleteBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error)

	GetAllBanners(ctx context.Context, callOptions ...callopt.Option) (resp []*product.BannerInfo, err error)

	AddNewBanner(ctx context.Context, req *product.NewBanner_, callOptions ...callopt.Option) (resp *product.BannerInfo, err error)

	UpdateBanner(ctx context.Context, req *product.BannerInfo, callOptions ...callopt.Option) (resp *product.BannerInfo, err error)

	DeleteBanner(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error)

	BatchAddCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_, callOptions ...callopt.Option) (resp []*product.CategoryBrandInfo, err error)

	BatchReduceCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_, callOptions ...callopt.Option) (resp *common.Empty, err error)

	DeleteBrandByCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error)

	DeleteCategoryByBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := combineservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient combineservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() combineservice.Client {
	return c.kitexClient
}

func (c *clientImpl) GetProductList(ctx context.Context, req *product.ProductFilter, callOptions ...callopt.Option) (resp []*product.ProductDetail, err error) {
	return c.kitexClient.GetProductList(ctx, req, callOptions...)
}

func (c *clientImpl) GetProductDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ProductDetail, err error) {
	return c.kitexClient.GetProductDetail(ctx, req, callOptions...)
}

func (c *clientImpl) AddNewProduct(ctx context.Context, req *product.NewProduct_, callOptions ...callopt.Option) (resp *product.ProductInfo, err error) {
	return c.kitexClient.AddNewProduct(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateProduct(ctx context.Context, req *product.ProductInfo, callOptions ...callopt.Option) (resp *product.ProductInfo, err error) {
	return c.kitexClient.UpdateProduct(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateRating(ctx context.Context, req *product.RatingReq, callOptions ...callopt.Option) (resp *product.RatingInfo, err error) {
	return c.kitexClient.UpdateRating(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteProduct(ctx, req, callOptions...)
}

func (c *clientImpl) GetCategoryList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (resp []*product.CategoryInfo, err error) {
	return c.kitexClient.GetCategoryList(ctx, req, callOptions...)
}

func (c *clientImpl) GetCategoryDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error) {
	return c.kitexClient.GetCategoryDetail(ctx, req, callOptions...)
}

func (c *clientImpl) AddNewCategory(ctx context.Context, req *product.NewCategory_, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error) {
	return c.kitexClient.AddNewCategory(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateCategory(ctx context.Context, req *product.CategoryInfo, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error) {
	return c.kitexClient.UpdateCategory(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteCategory(ctx, req, callOptions...)
}

func (c *clientImpl) GetAllModels(ctx context.Context, callOptions ...callopt.Option) (resp []*product.ModelInfo, err error) {
	return c.kitexClient.GetAllModels(ctx, callOptions...)
}

func (c *clientImpl) GetModelDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ModelInfo, err error) {
	return c.kitexClient.GetModelDetail(ctx, req, callOptions...)
}

func (c *clientImpl) AddNewModel(ctx context.Context, req *product.NewModel_, callOptions ...callopt.Option) (resp *product.ModelInfo, err error) {
	return c.kitexClient.AddNewModel(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteModel(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteModel(ctx, req, callOptions...)
}

func (c *clientImpl) GetBrandList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error) {
	return c.kitexClient.GetBrandList(ctx, req, callOptions...)
}

func (c *clientImpl) GetRelatedBrandsByCategoryId(ctx context.Context, req *product.BrandByCatReq, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error) {
	return c.kitexClient.GetRelatedBrandsByCategoryId(ctx, req, callOptions...)
}

func (c *clientImpl) GetBrandDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.BrandInfo, err error) {
	return c.kitexClient.GetBrandDetail(ctx, req, callOptions...)
}

func (c *clientImpl) AddNewBrand(ctx context.Context, req *product.NewBrand_, callOptions ...callopt.Option) (resp *product.BrandInfo, err error) {
	return c.kitexClient.AddNewBrand(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateBrand(ctx context.Context, req *product.BrandInfo, callOptions ...callopt.Option) (resp *product.BrandInfo, err error) {
	return c.kitexClient.UpdateBrand(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteBrand(ctx, req, callOptions...)
}

func (c *clientImpl) GetAllBanners(ctx context.Context, callOptions ...callopt.Option) (resp []*product.BannerInfo, err error) {
	return c.kitexClient.GetAllBanners(ctx, callOptions...)
}

func (c *clientImpl) AddNewBanner(ctx context.Context, req *product.NewBanner_, callOptions ...callopt.Option) (resp *product.BannerInfo, err error) {
	return c.kitexClient.AddNewBanner(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateBanner(ctx context.Context, req *product.BannerInfo, callOptions ...callopt.Option) (resp *product.BannerInfo, err error) {
	return c.kitexClient.UpdateBanner(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteBanner(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteBanner(ctx, req, callOptions...)
}

func (c *clientImpl) BatchAddCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_, callOptions ...callopt.Option) (resp []*product.CategoryBrandInfo, err error) {
	return c.kitexClient.BatchAddCategoryBrand(ctx, req, callOptions...)
}

func (c *clientImpl) BatchReduceCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.BatchReduceCategoryBrand(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteBrandByCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteBrandByCategory(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteCategoryByBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	return c.kitexClient.DeleteCategoryByBrand(ctx, req, callOptions...)
}
