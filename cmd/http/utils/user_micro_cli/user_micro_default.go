package user_micro_cli

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendSMSViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = defaultClient.SendSMSViaAliyun(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendSMSViaAliyun call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = defaultClient.FreezePatrolBeforeAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FreezePatrolBeforeAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FreezePatrolAfterAuth(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = defaultClient.FreezePatrolAfterAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FreezePatrolAfterAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyCredentials(ctx context.Context, req *micro_user.RpcCredentialReq, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = defaultClient.VerifyCredentials(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyCredentials call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcUser, err error) {
	resp, err = defaultClient.FindUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (resp []*micro_user.RpcUser, err error) {
	resp, err = defaultClient.QueryUsersWithFilter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryUsersWithFilter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpsertUser(ctx context.Context, req *micro_user.RpcUser, callOptions ...callopt.Option) (resp *micro_user.RpcUser, err error) {
	resp, err = defaultClient.UpsertUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpsertUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = defaultClient.DeleteUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
