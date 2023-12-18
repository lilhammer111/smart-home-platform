package micro_user_cli

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	micro_user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
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

func FreezePatrolAfterAuth(ctx context.Context, req *micro_user.RpcUserId, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = defaultClient.FreezePatrolAfterAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FreezePatrolAfterAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifySmsCode(ctx context.Context, req *micro_user.RpcVerifyCodeReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.VerifySmsCode(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifySmsCode call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyUsernamePwd(ctx context.Context, req *micro_user.RpcVerifyUsernamePwdReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.VerifyUsernamePwd(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyUsernamePwd call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyEmailPwd(ctx context.Context, req *micro_user.RpcVerifyEmailPwdReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.VerifyEmailPwd(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyEmailPwd call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindUser(ctx context.Context, req *micro_user.RpcFindUserReq, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.FindUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindUserByOpenid(ctx context.Context, req *micro_user.RpcFindUserByOpenidReq, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.FindUserByOpenid(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUserByOpenid call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindUserByMobile(ctx context.Context, req *micro_user.RpcFindUserByMobileReq, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.FindUserByMobile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUserByMobile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FindUserByUsername(ctx context.Context, username *micro_user.RpcFindUserByUsernameReq, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	resp, err = defaultClient.FindUserByUsername(ctx, username, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FindUserByUsername call failed,err =%+v", err)
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

func DeleteUser(ctx context.Context, userId int32, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.DeleteUser(ctx, userId, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
