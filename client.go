package aliyunmobile

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dypnsapi20170525 "github.com/alibabacloud-go/dypnsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// CreateClient 使用AK&SK初始化账号Client
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dypnsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("dypnsapi.aliyuncs.com")
	_result, _err = dypnsapi20170525.NewClient(config)
	return _result, _err
}

// GetMobileNumber 调用阿里云动态手机号保护服务获取手机号信息
// 修改为接受两个字符串参数：accessKeyID 和 accessKeySecret
func GetMobileNumber(accessKeyID, accessKeySecret string) (string, error) {
	client, err := CreateClient(tea.String(accessKeyID), tea.String(accessKeySecret))
	if err != nil {
		return "", fmt.Errorf("create client error: %v", err)
	}

	getMobileRequest := &dypnsapi20170525.GetMobileRequest{
		AccessToken: tea.String("your_access_token_here"), // 替换为实际的AccessToken
	}

	runtime := &util.RuntimeOptions{}

	// 发起获取号码请求
	response, err := client.GetMobileWithOptions(getMobileRequest, runtime)
	if err != nil {
		return "", fmt.Errorf("call GetMobileWithOptions error: %v", err)
	}

	// 处理返回的手机号信息
	if response.Body != nil && response.Body.Code == tea.String("OK") {
		return tea.StringValue(response.Body.GetMobileResultDTO.Mobile), nil
	} else if response.Body != nil {
		// 如果状态码不是OK，返回错误信息
		return "", fmt.Errorf("API error: %s - %s", tea.StringValue(response.Body.Code), tea.StringValue(response.Body.Message))
	}

	return "", fmt.Errorf("unexpected error: response body is nil")
}
