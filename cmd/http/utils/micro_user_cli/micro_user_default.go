package micro_user_cli

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = defaultClient.SendSmsViaAliyun(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendSmsViaAliyun call failed,err =%+v", err)
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

func FreezePatrolAfterAuth(ctx context.Context, id *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = defaultClient.FreezePatrolAfterAuth(ctx, id, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FreezePatrolAfterAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifySmsCode(ctx context.Context, mobile string, smsCode string, callOptions ...callopt.Option) (resp bool, err error) {
	resp, err = defaultClient.VerifySmsCode(ctx, mobile, smsCode, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifySmsCode call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyUsernamePwd(ctx context.Context, username string, entryPwd string, callOptions ...callopt.Option) (resp bool, err error) {
	resp, err = defaultClient.VerifyUsernamePwd(ctx, username, entryPwd, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyUsernamePwd call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyEmailPwd(ctx context.Context, email string, entryPwd string, callOptions ...callopt.Option) (resp bool, err error) {
	resp, err = defaultClient.VerifyEmailPwd(ctx, email, entryPwd, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyEmailPwd call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.FindUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (resp []*user.UserInfo, err error) {
	resp, err = defaultClient.QueryUsersWithFilter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryUsersWithFilter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.UpdateUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.CreateUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateUser call failed,err =%+v", err)
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
