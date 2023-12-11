package users

import (
	"context"
	users "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/kitex_gen/users"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FindUser(ctx context.Context, req int32, callOptions ...callopt.Option) (resp *users.UserData, err error) {
	resp, err = defaultClient.FindUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyCredentials(ctx context.Context, req *users.CredentialsReq, callOptions ...callopt.Option) (resp bool, err error) {
	resp, err = defaultClient.VerifyCredentials(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyCredentials call failed,err =%+v", err)
		return false, err
	}
	return resp, nil
}

func CreateOrUpdateUser(ctx context.Context, req *users.UserData, callOptions ...callopt.Option) (resp *users.UserData, err error) {
	resp, err = defaultClient.CreateOrUpdateUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOrUpdateUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req int32, callOptions ...callopt.Option) (resp *users.UserData, err error) {
	resp, err = defaultClient.DeleteUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
