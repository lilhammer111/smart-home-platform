// Code generated by Kitex v0.8.0. DO NOT EDIT.

package productservice

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetProductList(ctx context.Context, req *product.ProductFilter, callOptions ...callopt.Option) (r []*product.ProductDetail, err error)
	GetProductDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *product.ProductDetail, err error)
	AddNewProduct(ctx context.Context, req *product.NewProduct_, callOptions ...callopt.Option) (r *product.ProductInfo, err error)
	UpdateProduct(ctx context.Context, req *product.ProductInfo, callOptions ...callopt.Option) (r *product.ProductInfo, err error)
	UpdateRating(ctx context.Context, req *product.RatingReq, callOptions ...callopt.Option) (r *product.RatingInfo, err error)
	DeleteProduct(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductServiceClient struct {
	*kClient
}

func (p *kProductServiceClient) GetProductList(ctx context.Context, req *product.ProductFilter, callOptions ...callopt.Option) (r []*product.ProductDetail, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProductList(ctx, req)
}

func (p *kProductServiceClient) GetProductDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *product.ProductDetail, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProductDetail(ctx, req)
}

func (p *kProductServiceClient) AddNewProduct(ctx context.Context, req *product.NewProduct_, callOptions ...callopt.Option) (r *product.ProductInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNewProduct(ctx, req)
}

func (p *kProductServiceClient) UpdateProduct(ctx context.Context, req *product.ProductInfo, callOptions ...callopt.Option) (r *product.ProductInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateProduct(ctx, req)
}

func (p *kProductServiceClient) UpdateRating(ctx context.Context, req *product.RatingReq, callOptions ...callopt.Option) (r *product.RatingInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateRating(ctx, req)
}

func (p *kProductServiceClient) DeleteProduct(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteProduct(ctx, req)
}
