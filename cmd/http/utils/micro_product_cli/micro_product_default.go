package micro_product_cli

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetProductList(ctx context.Context, req *product.ProductFilter, callOptions ...callopt.Option) (resp []*product.BasicProdInfo, err error) {
	resp, err = defaultClient.GetProductList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *product.ProductInfo, err error) {
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

func UpdateProductBasicInfo(ctx context.Context, req *product.BasicProdInfo, callOptions ...callopt.Option) (resp *product.BasicProdInfo, err error) {
	resp, err = defaultClient.UpdateProductBasicInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProductBasicInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProdShowcase(ctx context.Context, req *product.ProductShowcase, callOptions ...callopt.Option) (resp *product.ProductShowcase, err error) {
	resp, err = defaultClient.UpdateProdShowcase(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProdShowcase call failed,err =%+v", err)
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

func GetBrandByCategory(ctx context.Context, req *product.GetBrandByCatReq, callOptions ...callopt.Option) (resp []*product.BrandInfo, err error) {
	resp, err = defaultClient.GetBrandByCategory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetBrandByCategory call failed,err =%+v", err)
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

func DelteBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DelteBrand(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DelteBrand call failed,err =%+v", err)
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

func DelteBanner(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DelteBanner(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DelteBanner call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
