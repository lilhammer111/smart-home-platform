package micro_product_cli

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetProductList(ctx context.Context, req *product.ProductFilter, callOptions ...callopt.Option) (resp []*product.ProductBasicInfo, err error) {
	resp, err = defaultClient.GetProductList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ProductDetailResp, err error) {
	resp, err = defaultClient.GetProductDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddNewProduct(ctx context.Context, req *product.AddProductReq, callOptions ...callopt.Option) (resp *product.ProductInfo, err error) {
	resp, err = defaultClient.AddNewProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddNewProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProduct(ctx context.Context, req *product.ProductInfo, callOptions ...callopt.Option) (resp *product.ProductInfo, err error) {
	resp, err = defaultClient.UpdateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RateProduct(ctx context.Context, req *product.RatingReq, callOptions ...callopt.Option) (resp *product.RatingResp, err error) {
	resp, err = defaultClient.RateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProduct(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCategoryList(ctx context.Context, req *common.PageFilter, callOptions ...callopt.Option) (resp []*product.CategoryInfo, err error) {
	resp, err = defaultClient.GetCategoryList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategoryList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCategoryDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error) {
	resp, err = defaultClient.GetCategoryDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategoryDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddNewCategory(ctx context.Context, req *product.NewCategory_, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error) {
	resp, err = defaultClient.AddNewCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddNewCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateCategory(ctx context.Context, req *product.CategoryInfo, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error) {
	resp, err = defaultClient.UpdateCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetBrandList(ctx context.Context, req *common.PageFilter, callOptions ...callopt.Option) (resp []*product.BrandListResp, err error) {
	resp, err = defaultClient.GetBrandList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetBrandList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetRelatedBrandsByCategoryId(ctx context.Context, req *product.BrandByCatReq, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error) {
	resp, err = defaultClient.GetRelatedBrandsByCategoryId(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetRelatedBrandsByCategoryId call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetBrandDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.BrandInfo, err error) {
	resp, err = defaultClient.GetBrandDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetBrandDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddNewBrand(ctx context.Context, req *product.NewBrand_, callOptions ...callopt.Option) (resp *product.BrandInfo, err error) {
	resp, err = defaultClient.AddNewBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddNewBrand call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateBrand(ctx context.Context, req *product.BrandInfo, callOptions ...callopt.Option) (resp *product.BrandInfo, err error) {
	resp, err = defaultClient.UpdateBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateBrand call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteBrand call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetAllBanners(ctx context.Context, callOptions ...callopt.Option) (resp []*product.BannerInfo, err error) {
	resp, err = defaultClient.GetAllBanners(ctx, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetAllBanners call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddNewBanner(ctx context.Context, req *product.NewBanner_, callOptions ...callopt.Option) (resp *product.BannerInfo, err error) {
	resp, err = defaultClient.AddNewBanner(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddNewBanner call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateBanner(ctx context.Context, req *product.BannerInfo, callOptions ...callopt.Option) (resp *product.BannerInfo, err error) {
	resp, err = defaultClient.UpdateBanner(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateBanner call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteBanner(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteBanner(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteBanner call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func BatchAddCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_, callOptions ...callopt.Option) (resp []*product.CategoryBrandInfo, err error) {
	resp, err = defaultClient.BatchAddCategoryBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BatchAddCategoryBrand call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func BatchReduceCategoryBrand(ctx context.Context, req *product.NewCategoryBrand_, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.BatchReduceCategoryBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "BatchReduceCategoryBrand call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteBrandByCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteBrandByCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteBrandByCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteCategoryByBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteCategoryByBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteCategoryByBrand call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
