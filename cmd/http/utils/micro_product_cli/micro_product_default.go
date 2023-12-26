package micro_product_cli

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetProductList(ctx context.Context, req *product.ProductFilter, callOptions ...callopt.Option) (resp []*product.ProductDetail, err error) {
	resp, err = defaultClient.GetProductList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ProductDetail, err error) {
	resp, err = defaultClient.GetProductDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddNewProduct(ctx context.Context, req *product.NewProduct_, callOptions ...callopt.Option) (resp *product.ProductInfo, err error) {
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

func UpdateRating(ctx context.Context, req *product.RatingReq, callOptions ...callopt.Option) (resp *product.RatingResp, err error) {
	resp, err = defaultClient.UpdateRating(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateRating call failed,err =%+v", err)
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

func GetCategoryList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (resp []*product.CategoryBasicInfo, err error) {
	resp, err = defaultClient.GetCategoryList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategoryList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCategoryDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.CategoryDetail, err error) {
	resp, err = defaultClient.GetCategoryDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategoryDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddNewCategory(ctx context.Context, req *product.AddCategoryReq, callOptions ...callopt.Option) (resp *product.CategoryInfo, err error) {
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

func GetAllModels(ctx context.Context, callOptions ...callopt.Option) (resp []*product.ModelInfo, err error) {
	resp, err = defaultClient.GetAllModels(ctx, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetAllModels call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetModelDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ModelInfo, err error) {
	resp, err = defaultClient.GetModelDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetModelDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddNewModel(ctx context.Context, req *product.NewModel_, callOptions ...callopt.Option) (resp *product.ModelInfo, err error) {
	resp, err = defaultClient.AddNewModel(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddNewModel call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteModel(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteModel(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteModel call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetBrandList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error) {
	resp, err = defaultClient.GetBrandList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetBrandList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetBrandListByCategory(ctx context.Context, req *product.BrandByCatReq, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error) {
	resp, err = defaultClient.GetBrandListByCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetBrandListByCategory call failed,err =%+v", err)
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

func GetCategoryBrandList(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp []*product.CategoryBrandInfo, err error) {
	resp, err = defaultClient.GetCategoryBrandList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCategoryBrandList call failed,err =%+v", err)
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

func UpdateCategoryBrand(ctx context.Context, req *product.CategoryBrandInfo, callOptions ...callopt.Option) (resp *product.CategoryBrandInfo, err error) {
	resp, err = defaultClient.UpdateCategoryBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateCategoryBrand call failed,err =%+v", err)
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

func DeleteBrandByCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteBrandByCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteBrandByCategory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
