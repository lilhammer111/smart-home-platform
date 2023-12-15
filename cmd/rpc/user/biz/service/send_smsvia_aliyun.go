package service

import (
	"context"
	"fmt"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/binding"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cloudwego/kitex/pkg/klog"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	accessKeyId     = "ACCESS_KEY_ID"
	accessKeySecret = "ACCESS_KEY_SECRET"
	endPoint        = "dysmsapi.aliyuncs.com"
	signName        = "阿里云短信测试"
	templateCode    = "SMS_154950909"

	smsTTL = 5 * time.Minute
)

type SendSMSViaAliyunService struct {
	ctx context.Context
}

// NewSendSMSViaAliyunService new SendSMSViaAliyunService
func NewSendSMSViaAliyunService(ctx context.Context) *SendSMSViaAliyunService {
	return &SendSMSViaAliyunService{ctx: ctx}
}

// Run create note info
func (s *SendSMSViaAliyunService) Run(req *micro_user.RpcSmsReq) (resp *common_rpc.RpcEmpty, err error) {
	// Finish your business logic.
	smsCode := generateSmsCode(6)

	smsSender := &SmsSender{}
	err = smsSender.initSmsAssistant(req.Mobile, smsCode)
	if err != nil {
		klog.Errorf("failed to get sms client: %s", err)
		return
	}

	smsResp, err := smsSender.smsAssistant.SendSms(smsSender.request)
	if err != nil || *smsResp.StatusCode != http.StatusOK {
		klog.Errorf("failed to send sms: %s", err)
		return
	}

	_, err = db.GetRedis().Set(s.ctx, req.Mobile, smsCode, smsTTL).Result()
	if err != nil {
		klog.Errorf("failed to set sms code in redis for mobile %s: %s", req.Mobile, err)
		return
	}
	return &common_rpc.RpcEmpty{}, err
}

// GenerateSmsCode helps to generate an SMS verification code with specified length.
func generateSmsCode(length int) string {
	numeric := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	var sb strings.Builder
	for i := 0; i < length; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			klog.Errorf("failed to generate a random number: %s", err)
		}
	}
	return sb.String()
}

type SmsSender struct {
	smsAssistant *dysmsapi.Client
	request      *dysmsapi.SendSmsRequest
}

func (ss *SmsSender) initSmsAssistant(mobile, smsCode string) (err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(binding.LoadEnv().GetEnvVar(accessKeyId)),
		AccessKeySecret: tea.String(binding.LoadEnv().GetEnvVar(accessKeySecret)),
		Endpoint:        tea.String(endPoint),
	}
	request := new(dysmsapi.SendSmsRequest)
	request.SetPhoneNumbers(mobile)
	request.SetSignName(signName)
	request.SetTemplateCode(templateCode)
	request.SetTemplateParam(fmt.Sprintf("{\"code\":%s}", smsCode))
	ss.smsAssistant, err = dysmsapi.NewClient(config)
	return
}
