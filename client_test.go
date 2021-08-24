package alipay

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigrocs/alipay/requests"
	"github.com/bigrocs/alipay/util"
)

func TestPay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.trade.pay"
	request.BizContent = map[string]interface{}{
		"subject":         "测试商品名称1",
		"scene":           "bar_code",
		"auth_code":       "280528574232947539",
		"out_trade_no":    "GZ2020010117534314525",
		"total_amount":    0.02,
		"timeout_express": "2m",
		"extend_params":   map[string]interface{}{"sys_service_provider_id": os.Getenv("PAY_ALIPAY_SYS_SERVICE_PROVIDER_ID")},
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	// fmt.Println("TestPay______", req, err)
	t.Log(req, err, "|||")
}

func TestPayQuery(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.trade.query"
	request.BizContent = map[string]interface{}{
		"out_trade_no": "GZ2020010117534314525",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	// fmt.Println("TestPayQuery______", req, err)
	t.Log(req, err, "|||")
}

func TestPayRefund(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.trade.refund"
	request.BizContent = map[string]interface{}{
		"out_trade_no":   "GZ2020010117534314525",
		"out_request_no": "GZ2020010117534314525_T2",
		"refund_amount":  0.01,
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	// fmt.Println("TestPayRefund______", req, err)
	t.Log(req, err, "|||")
}

func TestPayRefundQuery(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.trade.fastpay.refund.query"
	request.BizContent = map[string]interface{}{
		"out_trade_no":   "GZ2020010117534314525",
		"out_request_no": "GZ2020010117534314525_T21",
		"query_options": []string{
			"refund_royaltys",
		},
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	// fmt.Println("TestPayRefundQuery______", req, err)
	t.Log(req, err, "|||")
}

func TestQueryUserID(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.user.twostage.common.use"
	request.BizContent = map[string]interface{}{
		"sence_no":   "20170718101175343",
		"dynamic_id": "286861260475412123",
		"pay_pid":    "2088702093900999",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	// fmt.Println("TestQueryUserID______", req, err)
	t.Log(req, err, "|||")
}

// map[
// 	alipay_trade_pay_response:
// 	map[
// 		buyer_logon_id:rsk***@sandbox.com
// 		buyer_pay_amount:0.02
// 		buyer_user_id:2088102174959923
// 		buyer_user_type:PRIVATE
// 		code:10000
// 		fund_bill_list:[
// 			map[
// 				amount:0.02
// 				fund_channel:ALIPAYACCOUNT
// 			]
// 		]
// 		gmt_payment:2020-04-24 11:27:26
// 		invoice_amount:0.02
// 		msg:Success
// 		out_trade_no:GZ2020010117534314513
// 		point_amount:0.00
// 		receipt_amount:0.02
// 		total_amount:0.02
// 		trade_no:2020042422001459920500905861
// 	]
// 	sign:nq2l2bW72xitQYtCHn4XbOan1Lmwqhy+o0mO6cxzEVf4gJ4x6rm8RCiR9YeOeUl6aOygRFV+6I21t/fRsAK9LvNwtJMpaxyBY0yuQ60ycck17b5hxkWgYYjtYmFahfgnIo3UrZE/y0iLfBu2g/6P1Mp83g1MVU0gzIr1FbsBMACe9JastpNA/ggyZjRJA/fboEbNJAsI+QypB/7OvOXElsI7ZyGv1O6AkpUr01EjCILUQBSQP4VCz0QlCBHQWjduAQsx0s1Pryso9hNiOIPxr8DOsC61cQVM+rjhWeagOXMKm7QkIt2svbgaEuOmv5fmvA7dLgXV+3650nMvCws0wQ==
// ]

func TestSpBlueseaactivityCreate(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "2021001153698959"
	client.Config.PrivateKey = "MIIEowIBAAKCAQEApqebd4jnauq5unhvQr7zq8Qasfmxmc48QUU13tB6t7LIIAyit9Q0mgQz92J2BwfB8nX+1oByT81kELOtzYEq+a0ICbgEF4qUw4bZlWEU+lzlpn/pTCAVIEaoUIlXCrkE7osREKee0u2rjcU+jD9OqKEYXpa0YmsNs+R3XtL/CCGOAfkXTQAK8HKOR8itByKXmfwrhV6SML2CS+3CBno+4kIQfKm0MwTJndZPOiUp+IWDh2h4+9hFzrxVdWQez2IaRyMWgpR4Hla2CJr61c399JlOwh1lo49lSQNbZsy43GlXdo74NMSXUc8Ukn2N/1YFi1smNeN23y5mbzVlT+9nlQIDAQABAoIBABXrfbtfOnEJppwsMxYoX4x8Ira+DOoaAjrji45fBYUb8s6kWtKrCrAnVV54iOWr04q8kmRMSLzOWxcHSxj4oFpus9vzZDZkd+Sv2/aV6lhz15DqyhjFL4IrSQAu0FAMveVSTVBPe8hqkmH3S57OxN3RMo9eayN2yTqcFkdWoUzERrYKzyGkRsyATmCBjsLHJFkfcdgid0pjjXwgx24ICo9RfRzE5oS+dyY/KjPJCBrQrhTVDsZ2gkd8eIimqdA05Bq9lfP0TXMBBVpcPG/vlxGklUqL0EaITfTWwUecOW15kMabG6coEZqrt/cV/aDl0w8Zy82ydEmUFtpwphqlNRECgYEA1kcxZOAQRyAoAUHnlLrlwf19i5pYNHcmXUu7k0y6pV0WuH4xZzzIlBK3ElIiIMmeEfA1g1YtjMjDEw778ADRDMFKJkfb0miEg157cQJfYNy8Hkcer403cRWE8AzPyudifCPavDEYbvXB/kCqvux5xZ3e0l8yJgtWkyFSr/ga5rcCgYEAxxqZ7TcFnB0O8ylIHFOqROm3cEtbcHgM/c7ywoo/KpYJWyJnkmg/a3RdHbHdYbBdndsOQ5cEY8oFI1ySKZ0RdZjN3riTTfFDZcOd2gGkVn5GPAd+r9cg0beSMedtKJc1tpxcBcb6LrFsoz6FVZS3yAmdRA3dVEWEa2dSehYJ+BMCgYBB8oNHMpaP7Vuil2/4nFTeEXA0KPqAmzMZf5dpTDIddrrpaR03UsPAnsntP431XVfi2XT+yOUX8SnokbIEOZwdOtfhVDhZV4WBEPEY5MwJ+aWEBAtzqBiqGx9g6fCB2Bi9cCN/y/RjIgvkfm6OeCU6Js5PHLFIeDDa82SJd+h/nQKBgQCbE6ASgEmRR5zOTjgJQC0dkcA76aXl2suboGW6mn0KiWPuJMTw5tuGsVKZf2DD5t9zxRu+D1Du/Hm2CjMZ0Gm543Nt5utsyN7K02sOUzsR1zaTKO9GvFOmpMlsryl5dRT0D6MpKMrSNYU9PMQTXUsrtjlUuUvLiH3UPdyilKndGwKBgGXlAX5kNspdArdXHAgOiqiBvqEViCx9+meFlp0a7FhsDbXONslYLDQBWEirQIQbCguy9A0kwKr+O/ez8VpnsQu4n5dtRs6yjtvgJHU7JTFLrWw3vgx8sB+rXvDsHUXzizpGJq2FDlEjXfk4bjZD6f1r728A5HV70MbDA5BM9+ed"
	client.Config.AliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv0GV3fmf5g1ipb7HquzDUkZDxjE0cL4ZAZNj1YKncOFtAkp9GREYhwMlke8wkJK7XaIcDyCpkwPjMGw6m1qQv2dlmxF3h66tF0odTnuWoF/n0Neybs64IhIaK49Xys/B+ekS7kMOg7RYlGAr7lfwkRAKvq92ej80SXjWFv6K5B2YdhzTNH+86omhYzLakx8KGIZ8n3f7kMQXnZrMTm6gwK1FeTdxwEgiQrF3xO/O7hAuroElRoUBTbVnlDlXx0qXSl9cDDShMa0inkSUX4VSllFKxPwFtu6Nqsfj9O1N45xtPNwvHrpOWdYPzidvrGS1/MU+1sZKktPsNPYNiL58BQIDAQAB"
	// client.Config.Sandbox = true
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.open.sp.blueseaactivity.create"
	request.BizContent = map[string]interface{}{
		"biz_scene":      "BLUE_SEA_FMCG_APPLY",
		"merchant_logon": "xsjcssm@163.com",
		"province_code":  "370000",
		"city_code":      "371600",
		"district_code":  "371625",
		"address":        "博城五路新世纪超市",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	if req["sign"] != nil {
		ok, err := util.VerifySign(response.GetSignData(), req["sign"].(string), client.Config.AliPayPublicKey, "RSA2")
		fmt.Println(req, ok, err)
	}
	fmt.Println(req, err)
	t.Log(response, err)
}
