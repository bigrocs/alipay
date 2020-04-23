package alipay

import (
	"fmt"
	"testing"

	"github.com/bigrocs/alipay/requests"
)

func TestCode2Session(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wxa4153f8f312d3r8f7"

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "miniprogram"
	request.ApiName = "auth.code2Session"
	request.QueryParams = map[string]interface{}{
		"js_code": "0211MY2j255Vo1B0WKO3j2cqJ2j21MY2s",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	t.Log(response, err)
	t.Log(req, err)
}

func TestMchPayMicropay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "wxa4153f8f32d13r8f7"

	// 配置参数
	request := requests.NewCommonRequest()
	request.Domain = "mch"
	request.ApiName = "pay.micropay"
	request.QueryParams = map[string]interface{}{
		"auth_code":        "134770030978364234",
		"body":             "测试商品名称1",
		"out_trade_no":     "202002100007",
		"total_fee":        "1",
		"spbill_create_ip": "127.0.0.1",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	// fmt.Println(response, err)
	fmt.Println(req, err)
}
