// Code generated by Kitex v0.8.0. DO NOT EDIT.

package categoryservice

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetCategoryList(ctx context.Context, req *common.PageFilter, callOptions ...callopt.Option) (r []*product.CategoryInfo, err error)
	GetCategoryDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *product.CategoryInfo, err error)
	AddNewCategory(ctx context.Context, req *product.NewCategory_, callOptions ...callopt.Option) (r *product.CategoryInfo, err error)
	UpdateCategory(ctx context.Context, req *product.CategoryInfo, callOptions ...callopt.Option) (r *product.CategoryInfo, err error)
	DeleteCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error)
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
	return &kCategoryServiceClient{
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

type kCategoryServiceClient struct {
	*kClient
}

func (p *kCategoryServiceClient) GetCategoryList(ctx context.Context, req *common.PageFilter, callOptions ...callopt.Option) (r []*product.CategoryInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCategoryList(ctx, req)
}

func (p *kCategoryServiceClient) GetCategoryDetail(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *product.CategoryInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCategoryDetail(ctx, req)
}

func (p *kCategoryServiceClient) AddNewCategory(ctx context.Context, req *product.NewCategory_, callOptions ...callopt.Option) (r *product.CategoryInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNewCategory(ctx, req)
}

func (p *kCategoryServiceClient) UpdateCategory(ctx context.Context, req *product.CategoryInfo, callOptions ...callopt.Option) (r *product.CategoryInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCategory(ctx, req)
}

func (p *kCategoryServiceClient) DeleteCategory(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCategory(ctx, req)
}
