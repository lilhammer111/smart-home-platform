// Code generated by Kitex v0.6.1. DO NOT EDIT.

package modelservice

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetAllModels(ctx context.Context, callOptions ...callopt.Option) (r []*product.ModelInfo, err error)
	AddNewModel(ctx context.Context, req *product.NewModel_, callOptions ...callopt.Option) (r *product.ModelInfo, err error)
	DeleteModel(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error)
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
	return &kModelServiceClient{
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

type kModelServiceClient struct {
	*kClient
}

func (p *kModelServiceClient) GetAllModels(ctx context.Context, callOptions ...callopt.Option) (r []*product.ModelInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllModels(ctx)
}

func (p *kModelServiceClient) AddNewModel(ctx context.Context, req *product.NewModel_, callOptions ...callopt.Option) (r *product.ModelInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNewModel(ctx, req)
}

func (p *kModelServiceClient) DeleteModel(ctx context.Context, req *common.Req, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteModel(ctx, req)
}
