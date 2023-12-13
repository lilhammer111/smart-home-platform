package user_micro

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	micro_user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.FreezeRpcReq, callOptions ...callopt.Option) (resp *micro_user.FreezeRpcResp, err error) {
	resp, err = defaultClient.FreezePatrolBeforeAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FreezePatrolBeforeAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FreezePatrolAfterAuth(ctx context.Context, req *common.IdRpcReq, callOptions ...callopt.Option) (resp *micro_user.FreezeRpcResp, err error) {
	resp, err = defaultClient.FreezePatrolAfterAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FreezePatrolAfterAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyCredentials(ctx context.Context, req *micro_user.CredentialRpcReq, callOptions ...callopt.Option) (resp *common.EmptyRpcResp, err error) {
	resp, err = defaultClient.VerifyCredentials(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyCredentials call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindUser(ctx context.Context, req *common.IdRpcReq, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.FindUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryUsersWithFilter(ctx context.Context, req *user.UsersFilter, callOptions ...callopt.Option) (resp []*user.UserInfo, err error) {
	resp, err = defaultClient.QueryUsersWithFilter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryUsersWithFilter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpsertUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.UpsertUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpsertUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *common.IdRpcReq, callOptions ...callopt.Option) (resp *common.EmptyRpcResp, err error) {
	resp, err = defaultClient.DeleteUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
