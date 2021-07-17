package alipay

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigrocs/alipay/requests"
)

func TestPay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = os.Getenv("ALIPAY_APPID")
	client.Config.PrivateKey = os.Getenv("ALIPAY_PRIVATE_KEY")
	client.Config.AliPayPublicKey = os.Getenv("ALIPAY_PUBLIC_KEY")
	client.Config.AppAuthToken = os.Getenv("ALIPAY_APP_AUTH_TOKEN")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "alipay.trade.pay"
	request.BizContent = map[string]interface{}{
		"subject":         "测试商品名称1",
		"scene":           "bar_code",
		"auth_code":       "282570666284142748",
		"out_trade_no":    "GZ2020010117534314522",
		"total_amount":    100000,
		"timeout_express": "2m",
		"extend_params":   map[string]interface{}{"sys_service_provider_id": os.Getenv("ALIPAY_SYS_SERVICE_PROVIDER_ID")},
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	r, err := response.GetVerifySignDataMap()
	fmt.Println("TestPay______", r, err)
	t.Log(req, err, "|||")
}

// func TestPayQuery(t *testing.T) {
// 	// 创建连接
// 	client := NewClient()
// 	client.Config.AppId = "2016080900197401"
// 	client.Config.Method = "alipay.trade.query"
// 	client.Config.PrivateKey = "MIIEogIBAAKCAQEAgc4wW8rlV9XvywSUJYsib7Csg6waonzGAISqZjDNqcDY2oYK/yTwHd9L1qH69fS6sTDY32R0oSFy9Pg22f040GAALECRTtcGkyBC1wsfQUICxhTASbOX1NzY4Hk//8+4dcNAGAVxeByLx52hKJyCMjDveTwZLRCijxsJ56kgcjulptA5LxsiImInxHyH5Ne+kgPcZUt4o5IbpcngpNZ2GVrASbPzV/5BiuBza883ey1c4Ps/dHjA7vNOjfJ/lpIxWLdmfrs1aHTkZdRsRcsvPWzBIYZ1mpUIlGF48S1hkptwA8BpNCR2Lhn1NO8fQkZkDxHxhBp5IFtp92bJ+PJ1iwIDAQABAoIBAAKhWvmBMwSEoUdL4HSiTSBWRCim5CNGw/xes9U+R+yYq8xByxdAeF0imTbQMXWqb94HT123e2WId/vey72/E9elGlvLMyaV1NxGWxRgdVGtNMEuynaz3O/DSeHCkzlrzCEWw8W41oPIkEyLNSjNgxXhR/j87sWrAKBFu/2lv7KQ23Me6rY5hiJq7RW+5z0deEqDLkPHmbM12lJoNAvmhX4Q715s1LivUsFP5KZkxKzGgh233S2BcGuTIYj036g1DI83GbG1cVSDW9gU+J43CKKOYu+dCqUpYBXvRoTFDySqTo6blv4iQD/rvDvqFJkNBsEHV2HNbLNK/fREpK43IwECgYEAwqRwjp5+osbYAtHSz+GEFoWgTozKSxYHmYc55xRvi9bE328Y4ZxyyXZiQCxFKwdmCQl/E2sm1MalRwKRRhDRnIgkIIS/wGuwJ7fWkeNjmO68MaExs/mmdn9eYQJA8knms/sB19qAFuPl9uSiY5CJfLxaTqqTzF1fM02ot3uM/7MCgYEAqrlvgXo0VPWNUxVdMRHXPaRk8kl+lFYb8qXJVU2X/vdShHJ4iPVeFUI/XTeOFWherHkH83u1BWhkLU+KvrYxyG0rxdINYI/priMKDjQ06tH2cQshbe1D66ku9TnM1kbiYN6xJpu3MJsre0v/hqgr+jQ9ucMCVkMLZRIb9i4ChskCgYBWJxLrDZxf0EOse9Mj0F322g6qtgaUVZvniMNIVPNKojVh8HrC4cys/4ldfjrfYNb5CQsGPVditspRNAG5UZh2AIx9GEUHlqLR8b03tb5P4tvJ2990GfxkVtwfdC/rDmrfHyshr8UiXJ1dOrXl/APfAj+2sinZOzr4KleTX0x5oQKBgFFEC8v6O8blS+xskvu0nlx9UH/0dAhwJMWQHRI56Uw4tlLNmoq3IZ3E9xyMQVn3YHmA80P3cuesFWNsJYM6fuAE930my69XUcjObQ7t0vKkF0cgIT2OX3JiCjQ420R0YXXzCyhxnvXIJx59wph6nNRw4aD3LrmZGGd4A09f/1vhAoGANc3hlyUk28qctgQQSi10/Ng5Fy4YIKW+g8WwNZ3voaeIyaSM/9jJ+teqqnJw8mcgGYo+YMqU+7IEwzrcjib4id2cu4MhlRswXKgaG2gSIUN7JEivjQisM6UUkXZtD91ZKx9JxTPVRYXqAjCcd6YPYt/C8PJTfGpQUFttX498P5g="
// 	client.Config.AliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAruH561raPfR7mk3OFe/AKGsuir0oqfnwRekRtAo4EbliOlfSB2XAcQnyn6Wkc9bvRWkgGq6MJV2lVKRs114yyGz1MEhrjz8P1slp3KFnx/TwQgZSTGVH55BLNfB0cc+YA7/beTXHCOG4rQp8KPLURplkCMtuM/dQwS/6b/pF6dFHFhkZgXsHwtzK20jr6xVcT2Hk4tQGA1tfUSrskkj+CH61TSGfp5YkkfnieG3FEGfCjod0t37dCDKFNxD6EDOa10VqFtipLspo14PTDQmr3wQHCfZfmXqMdHtr2NMnIDYT4DCHhcUSI0VPMAohLbW6Y4Dm1JEkOyighLbrgY2qYQIDAQAB"
// 	client.Config.Sandbox = true
// 	// 配置参数
// 	request := requests.NewCommonRequest()
// 	request.BizContent = map[string]interface{}{
// 		"out_trade_no": "GZ2020010117534314513",
// 	}
// 	// 请求
// 	response, err := client.ProcessCommonRequest(request)
// 	req, err := response.GetHttpContentMap()
// 	if req["sign"] != nil {
// 		ok, err := util.VerifySign(response.GetSignData(), req["sign"].(string), client.Config.AliPayPublicKey, "RSA2")
// 		fmt.Println(req, ok, err)
// 	}
// 	t.Log(response, err)
// }

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
	// client.Config.Method = "alipay.open.sp.blueseaactivity.create"
	// client.Config.PrivateKey = "MIIEowIBAAKCAQEApqebd4jnauq5unhvQr7zq8Qasfmxmc48QUU13tB6t7LIIAyit9Q0mgQz92J2BwfB8nX+1oByT81kELOtzYEq+a0ICbgEF4qUw4bZlWEU+lzlpn/pTCAVIEaoUIlXCrkE7osREKee0u2rjcU+jD9OqKEYXpa0YmsNs+R3XtL/CCGOAfkXTQAK8HKOR8itByKXmfwrhV6SML2CS+3CBno+4kIQfKm0MwTJndZPOiUp+IWDh2h4+9hFzrxVdWQez2IaRyMWgpR4Hla2CJr61c399JlOwh1lo49lSQNbZsy43GlXdo74NMSXUc8Ukn2N/1YFi1smNeN23y5mbzVlT+9nlQIDAQABAoIBABXrfbtfOnEJppwsMxYoX4x8Ira+DOoaAjrji45fBYUb8s6kWtKrCrAnVV54iOWr04q8kmRMSLzOWxcHSxj4oFpus9vzZDZkd+Sv2/aV6lhz15DqyhjFL4IrSQAu0FAMveVSTVBPe8hqkmH3S57OxN3RMo9eayN2yTqcFkdWoUzERrYKzyGkRsyATmCBjsLHJFkfcdgid0pjjXwgx24ICo9RfRzE5oS+dyY/KjPJCBrQrhTVDsZ2gkd8eIimqdA05Bq9lfP0TXMBBVpcPG/vlxGklUqL0EaITfTWwUecOW15kMabG6coEZqrt/cV/aDl0w8Zy82ydEmUFtpwphqlNRECgYEA1kcxZOAQRyAoAUHnlLrlwf19i5pYNHcmXUu7k0y6pV0WuH4xZzzIlBK3ElIiIMmeEfA1g1YtjMjDEw778ADRDMFKJkfb0miEg157cQJfYNy8Hkcer403cRWE8AzPyudifCPavDEYbvXB/kCqvux5xZ3e0l8yJgtWkyFSr/ga5rcCgYEAxxqZ7TcFnB0O8ylIHFOqROm3cEtbcHgM/c7ywoo/KpYJWyJnkmg/a3RdHbHdYbBdndsOQ5cEY8oFI1ySKZ0RdZjN3riTTfFDZcOd2gGkVn5GPAd+r9cg0beSMedtKJc1tpxcBcb6LrFsoz6FVZS3yAmdRA3dVEWEa2dSehYJ+BMCgYBB8oNHMpaP7Vuil2/4nFTeEXA0KPqAmzMZf5dpTDIddrrpaR03UsPAnsntP431XVfi2XT+yOUX8SnokbIEOZwdOtfhVDhZV4WBEPEY5MwJ+aWEBAtzqBiqGx9g6fCB2Bi9cCN/y/RjIgvkfm6OeCU6Js5PHLFIeDDa82SJd+h/nQKBgQCbE6ASgEmRR5zOTjgJQC0dkcA76aXl2suboGW6mn0KiWPuJMTw5tuGsVKZf2DD5t9zxRu+D1Du/Hm2CjMZ0Gm543Nt5utsyN7K02sOUzsR1zaTKO9GvFOmpMlsryl5dRT0D6MpKMrSNYU9PMQTXUsrtjlUuUvLiH3UPdyilKndGwKBgGXlAX5kNspdArdXHAgOiqiBvqEViCx9+meFlp0a7FhsDbXONslYLDQBWEirQIQbCguy9A0kwKr+O/ez8VpnsQu4n5dtRs6yjtvgJHU7JTFLrWw3vgx8sB+rXvDsHUXzizpGJq2FDlEjXfk4bjZD6f1r728A5HV70MbDA5BM9+ed"
	// client.Config.AliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv0GV3fmf5g1ipb7HquzDUkZDxjE0cL4ZAZNj1YKncOFtAkp9GREYhwMlke8wkJK7XaIcDyCpkwPjMGw6m1qQv2dlmxF3h66tF0odTnuWoF/n0Neybs64IhIaK49Xys/B+ekS7kMOg7RYlGAr7lfwkRAKvq92ej80SXjWFv6K5B2YdhzTNH+86omhYzLakx8KGIZ8n3f7kMQXnZrMTm6gwK1FeTdxwEgiQrF3xO/O7hAuroElRoUBTbVnlDlXx0qXSl9cDDShMa0inkSUX4VSllFKxPwFtu6Nqsfj9O1N45xtPNwvHrpOWdYPzidvrGS1/MU+1sZKktPsNPYNiL58BQIDAQAB"
	// // client.Config.Sandbox = true
	// // 配置参数
	// request := requests.NewCommonRequest()
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
	// t.Log(response, err)
}
