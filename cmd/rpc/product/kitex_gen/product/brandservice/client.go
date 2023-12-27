// Code generated by Kitex v0.8.0. DO NOT EDIT.

package brandservice

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetBrandList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (r []*product.BrandInfo, err error)
	GetRelatedBrandsByCategoryId(ctx context.Context, req *product.BrandByCatReq, callOptions ...callopt.Option) (r []*product.BrandInfo, err error)
	GetBrandDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *product.BrandInfo, err error)
	AddNewBrand(ctx context.Context, req *product.NewBrand_, callOptions ...callopt.Option) (r *product.BrandInfo, err error)
	UpdateBrand(ctx context.Context, req *product.BrandInfo, callOptions ...callopt.Option) (r *product.BrandInfo, err error)
	DeleteBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error)
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
	return &kBrandServiceClient{
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

type kBrandServiceClient struct {
	*kClient
}

func (p *kBrandServiceClient) GetBrandList(ctx context.Context, req *product.PageFilter, callOptions ...callopt.Option) (r []*product.BrandInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBrandList(ctx, req)
}

func (p *kBrandServiceClient) GetRelatedBrandsByCategoryId(ctx context.Context, req *product.BrandByCatReq, callOptions ...callopt.Option) (r []*product.BrandInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRelatedBrandsByCategoryId(ctx, req)
}

func (p *kBrandServiceClient) GetBrandDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *product.BrandInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBrandDetail(ctx, req)
}

func (p *kBrandServiceClient) AddNewBrand(ctx context.Context, req *product.NewBrand_, callOptions ...callopt.Option) (r *product.BrandInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNewBrand(ctx, req)
}

func (p *kBrandServiceClient) UpdateBrand(ctx context.Context, req *product.BrandInfo, callOptions ...callopt.Option) (r *product.BrandInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateBrand(ctx, req)
}

func (p *kBrandServiceClient) DeleteBrand(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteBrand(ctx, req)
}
