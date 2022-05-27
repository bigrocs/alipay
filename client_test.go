package alipay

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigrocs/alipay/requests"
)

func TestPay(t *testing.T) {
	// 创建连接
	// client := NewClient()
	// client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	// client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	// client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	// client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	// client.Config.Sandbox = false
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.ApiName = "alipay.trade.pay"
	// request.BizContent = map[string]interface{}{
	// 	"subject":         "测试商品名称1",
	// 	"scene":           "bar_code",
	// 	"auth_code":       "287643781357659050",
	// 	"out_trade_no":    "GZ2020010117534314526",
	// 	"total_amount":    0.01,
	// 	"timeout_express": "2m",
	// 	"extend_params":   map[string]interface{}{"sys_service_provider_id": os.Getenv("PAY_ALIPAY_SYS_SERVICE_PROVIDER_ID")},
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetVerifySignDataMap()
	// // fmt.Println("TestPay______", req, err)
	// t.Log(req, err, "|||")
}

func TestPayQuery(t *testing.T) {
	// 创建连接
	// client := NewClient()
	// client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	// client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	// client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	// client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	// client.Config.Sandbox = false
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.ApiName = "alipay.trade.query"
	// request.BizContent = map[string]interface{}{
	// 	"out_trade_no": "GZ2020010117534314525",
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetVerifySignDataMap()
	// // fmt.Println("TestPayQuery______", req, err)
	// t.Log(req, err, "|||")
}
func TestPayCreate(t *testing.T) {
	// // 创建连接
	// client := NewClient()
	// client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	// client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	// client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	// client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	// client.Config.Sandbox = false
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.ApiName = "alipay.trade.create"
	// request.BizContent = map[string]interface{}{
	// 	"subject":         "测试商品名称1",
	// 	"out_trade_no":    "GZ2020010117534314527",
	// 	"total_amount":    0.01,
	// 	"timeout_express": "2m",
	// 	"buyer_id":        "2088002104076813",
	// 	"extend_params":   map[string]interface{}{"sys_service_provider_id": os.Getenv("PAY_ALIPAY_SYS_SERVICE_PROVIDER_ID")},
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetVerifySignDataMap()
	// fmt.Println("TestPayQuery______", req, err)
	// t.Log(req, err, "|||")
}

func TestPayRefund(t *testing.T) {
	// 创建连接
	// client := NewClient()
	// client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	// client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	// client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	// client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	// client.Config.Sandbox = false
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.ApiName = "alipay.trade.refund"
	// request.BizContent = map[string]interface{}{
	// 	"out_trade_no":   "GZ2020010117534314526",
	// 	"out_request_no": "GZ2020010117534314526_1",
	// 	"refund_amount":  0.01,
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetVerifySignDataMap()
	// // fmt.Println("TestPayRefund______", req, err)
	// t.Log(req, err, "|||")
}

func TestPayRefundQuery(t *testing.T) {
	// 创建连接
	// client := NewClient()
	// client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	// client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	// client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	// client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	// client.Config.Sandbox = false
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.ApiName = "alipay.trade.fastpay.refund.query"
	// request.BizContent = map[string]interface{}{
	// 	"out_trade_no":   "GZ2020010117534314526",
	// 	"out_request_no": "GZ2020010117534314526_1",
	// 	// "query_options": []string{
	// 	// 	"refund_royaltys",
	// 	// },
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetVerifySignDataMap()
	// // fmt.Println("TestPayRefundQuery______", req, err)
	// t.Log(req, err, "|||")
}

func TestQueryUserID(t *testing.T) {
	// // 创建连接
	// client := NewClient()
	// client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	// client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	// client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	// client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	// client.Config.Sandbox = false
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.ApiName = "alipay.mobile.shake.user.query"
	// request.BizContent = map[string]interface{}{
	// 	"dynamic_id_type": "qr_code",
	// 	"dynamic_id":      "285479049183432686",
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetVerifySignDataMap()
	// // fmt.Println("TestQueryUserID______", req, err)
	// t.Log(req, err, "|||")
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
	// client := NewClient()
	// client.Config.AppId = "2021001153698959"
	// client.Config.PrivateKey = "MIIEowIBAAKCAQEApqebd4jnauq5unhvQr7zq8Qasfmxmc48QUU13tB6t7LIIAyit9Q0mgQz92J2BwfB8nX+1oByT81kELOtzYEq+a0ICbgEF4qUw4bZlWEU+lzlpn/pTCAVIEaoUIlXCrkE7osREKee0u2rjcU+jD9OqKEYXpa0YmsNs+R3XtL/CCGOAfkXTQAK8HKOR8itByKXmfwrhV6SML2CS+3CBno+4kIQfKm0MwTJndZPOiUp+IWDh2h4+9hFzrxVdWQez2IaRyMWgpR4Hla2CJr61c399JlOwh1lo49lSQNbZsy43GlXdo74NMSXUc8Ukn2N/1YFi1smNeN23y5mbzVlT+9nlQIDAQABAoIBABXrfbtfOnEJppwsMxYoX4x8Ira+DOoaAjrji45fBYUb8s6kWtKrCrAnVV54iOWr04q8kmRMSLzOWxcHSxj4oFpus9vzZDZkd+Sv2/aV6lhz15DqyhjFL4IrSQAu0FAMveVSTVBPe8hqkmH3S57OxN3RMo9eayN2yTqcFkdWoUzERrYKzyGkRsyATmCBjsLHJFkfcdgid0pjjXwgx24ICo9RfRzE5oS+dyY/KjPJCBrQrhTVDsZ2gkd8eIimqdA05Bq9lfP0TXMBBVpcPG/vlxGklUqL0EaITfTWwUecOW15kMabG6coEZqrt/cV/aDl0w8Zy82ydEmUFtpwphqlNRECgYEA1kcxZOAQRyAoAUHnlLrlwf19i5pYNHcmXUu7k0y6pV0WuH4xZzzIlBK3ElIiIMmeEfA1g1YtjMjDEw778ADRDMFKJkfb0miEg157cQJfYNy8Hkcer403cRWE8AzPyudifCPavDEYbvXB/kCqvux5xZ3e0l8yJgtWkyFSr/ga5rcCgYEAxxqZ7TcFnB0O8ylIHFOqROm3cEtbcHgM/c7ywoo/KpYJWyJnkmg/a3RdHbHdYbBdndsOQ5cEY8oFI1ySKZ0RdZjN3riTTfFDZcOd2gGkVn5GPAd+r9cg0beSMedtKJc1tpxcBcb6LrFsoz6FVZS3yAmdRA3dVEWEa2dSehYJ+BMCgYBB8oNHMpaP7Vuil2/4nFTeEXA0KPqAmzMZf5dpTDIddrrpaR03UsPAnsntP431XVfi2XT+yOUX8SnokbIEOZwdOtfhVDhZV4WBEPEY5MwJ+aWEBAtzqBiqGx9g6fCB2Bi9cCN/y/RjIgvkfm6OeCU6Js5PHLFIeDDa82SJd+h/nQKBgQCbE6ASgEmRR5zOTjgJQC0dkcA76aXl2suboGW6mn0KiWPuJMTw5tuGsVKZf2DD5t9zxRu+D1Du/Hm2CjMZ0Gm543Nt5utsyN7K02sOUzsR1zaTKO9GvFOmpMlsryl5dRT0D6MpKMrSNYU9PMQTXUsrtjlUuUvLiH3UPdyilKndGwKBgGXlAX5kNspdArdXHAgOiqiBvqEViCx9+meFlp0a7FhsDbXONslYLDQBWEirQIQbCguy9A0kwKr+O/ez8VpnsQu4n5dtRs6yjtvgJHU7JTFLrWw3vgx8sB+rXvDsHUXzizpGJq2FDlEjXfk4bjZD6f1r728A5HV70MbDA5BM9+ed"
	// client.Config.AliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv0GV3fmf5g1ipb7HquzDUkZDxjE0cL4ZAZNj1YKncOFtAkp9GREYhwMlke8wkJK7XaIcDyCpkwPjMGw6m1qQv2dlmxF3h66tF0odTnuWoF/n0Neybs64IhIaK49Xys/B+ekS7kMOg7RYlGAr7lfwkRAKvq92ej80SXjWFv6K5B2YdhzTNH+86omhYzLakx8KGIZ8n3f7kMQXnZrMTm6gwK1FeTdxwEgiQrF3xO/O7hAuroElRoUBTbVnlDlXx0qXSl9cDDShMa0inkSUX4VSllFKxPwFtu6Nqsfj9O1N45xtPNwvHrpOWdYPzidvrGS1/MU+1sZKktPsNPYNiL58BQIDAQAB"
	// // client.Config.Sandbox = true
	// // 配置参数
	// request := requests.NewCommonRequest()
	// request.ApiName = "alipay.open.sp.blueseaactivity.create"
	// request.BizContent = map[string]interface{}{
	// 	"biz_scene":      "BLUE_SEA_FMCG_APPLY",
	// 	"merchant_logon": "xsjcssm@163.com",
	// 	"province_code":  "370000",
	// 	"city_code":      "371600",
	// 	"district_code":  "371625",
	// 	"address":        "博城五路新世纪超市",
	// }
	// // 请求
	// response, err := client.ProcessCommonRequest(request)
	// req, err := response.GetHttpContentMap()
	// if req["sign"] != nil {
	// 	ok, err := util.VerifySign(response.GetSignData(), req["sign"].(string), client.Config.AliPayPublicKey, "RSA2")
	// 	fmt.Println(req, ok, err)
	// }
	// fmt.Println(req, err)
	// t.Log(response, err)
}

func TestOauthToken(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("PAY_ALIPAY_APPID")
	client.Config.PrivateKey = os.Getenv("PAY_ALIPAY_PRIVATE_KEY")
	client.Config.AliPayPublicKey = os.Getenv("PAY_ALIPAY_ALIPAY_ALIPAY_PUBLIC_KEY")
	// client.Config.AppAuthToken = os.Getenv("PAY_ALIPAY_APP_AUTH_TOKEN")
	client.Config.AppId = "2021002195639418"
	client.Config.PrivateKey = "MIIEpAIBAAKCAQEAg0EaudAK+QxMM+btSqGk9TGBP/xo0WANoxRto0n56Tn9IJWIMM3MuGPVZLemSpndC+tOA9SCRhSCyJ5DNukGXOUaE250Pfb3uWSNXsuyx33L8vb27tzbYG49fnOy0vAxYbDANDSslQ9q9eJ6gKbmqrVMc6ZdQzurqww99HOcr7gUTrGnrUvZK6pxdMrpcpBKbqcI1GeZdSJKkVl5kqRN8Ryxu/mFsgmhD1ALUPWQkD4mL8RQEs6EbngWdYucLTTRPXVhXxYu3KKOkE8HoxN5lIi4pecot2CT9PN4qoj43dK9B4WHEBIpIjVJo3xNQh3UtqEZKDwDxDE1m4M2uqRH+wIDAQABAoIBABFmYBSh79J0PuXihzKNm0Cl/2xsGbkn8qbb3mz3MZophdT7/cH6hlkrYu6nKfRZ0WPzHbSER4Di1zOhR2fQHHfUJ1DMI7zTMKPzIIqglmK72n2rH7EqtipicuSAkgREs0cwF0rUcMyR7qwKChNkDNA0mJ38QADZuDhEtzHg4bSpkdZ2qBsHjkOvXI6XPNa1yxbKzr9au/kbH4CxsA7bs5WOrdbrRayhH9IEjr2Yd0GaUYMGYe83aWq5rzr0vyqHhw3ZT1tKdUGLhYqpc5C5PeRPIJkkVfZLDGm5WOdZlyF+s31pApSnpA8dUKxvSk+w2JhgDDPXdA6YATXMsXHvrwECgYEAwtxImchqK1/OlxT4bxagLdedLSYGeiiAyVWFHGk7lhSbB4UpOxemggY8azwTfEaNTSn2oIThazxnSptFAbFsNWA2DFA20GR8NzNPejJ6tHbBazWfPCZOP12xboOcXbQV1HFBCFzS/n0af9v4lVy9sV9dvUE7+JBz0tEWpQQJmAkCgYEArG/NhC33kDwikGqNx7z65vHbBVdaMCiXWcE3WOcPEd9Ev0O+8q38643l3zFt9Vi7IMMYgD5yRAj1KJeu99pZTIBcLGCd0hfzNx9Eb+jpbfgmZEvVYw9kSKnySU5z9VsRn6BFcCdlGQCGwfLjskTAisoLHvhWFla+VaeZJNj8uOMCgYEAuFH2usDA7NZrbK4BcdNNIQ/bNWHKrrKxX8DtcLp6CrTjEHH9VlBAyK1taWIibGOwNVWEXMJVGELC0eedlsRlPvgLKAqYTvc+KFj4zPYC68GJpAfM3/C8A2R0+by0pE48d//JLEfhrahY8Kj6sw5OwQEcgnC6cZzycnAlo2H2UUECgYEAiEDbywsraDwsL6MSnvZz51LKduD8MbuRRGnJWyVQohWi0+MrGcVjaZm5PvBm64wOaSffz04IqjvNxCJ/LffJwgpXSvWZHy6vjynUwPzwnmWLqIrHdTvvSP+IpGYqeIBaX6Cb3bAzUVs6IDzVsMJkxgY6XWFrXMF0KmiBOT2ADHMCgYBV1P3Yg/7rGrASD2z5xTG9JI8e8AWgPp6HNccb56y9Ioc6RzSFkMSJ9/cZkfMoZTv6W8jCbvhztEnXtvzU7MAMQTuhpuTr3Ze3tc3ddb1q2Ay5Z435cbnILmMsGEy7JL6BIMo5Lztkr7vd1CCZk+m+SfEEBbO9GUZE8rnb+dSdsA=="
	client.Config.AliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAn85OVdsGHw8nGxuc/pS5D28fP2XfPesMatMBajlB3Yxsmz36JA95tGwlmrdhN2c1Cug/FfoWeayDMV35qjYhV/5N7UzW/AF6gfR/4UWrmCtfhU8PCocZpmTRq+EfNuzNiuwCJ/jApDiYDtVJGxMzs17YbX1k6+pTk05rHVj8jLNpMfHOlxKfXJHtpxN6l+pUYfX/XCFjyuiOYCU+x+VvnsudLmwPhmTpyoLeDNCiXGd8V8rTn5uffxDuptQtiu1V4H6mF5gAYbXbH9Oq6X2zA2BXEXLHwo+kJOS3vMQumVIWD2eFkGzxIXQHg7oyRd7qEfRei1KzwsC7PaavnNQhmQIDAQAB"
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.system.oauth.token"
	request.BizContent = map[string]interface{}{
		"code":       "9a6d0d639d3c475e99c1132dc2d5UX81",
		"grant_type": "authorization_code",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetVerifySignDataMap()
	fmt.Println("TestOauthToken______", req, err)
	t.Log(req, err, "|||")
}
