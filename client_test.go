package alipay

import (
	"fmt"
	"testing"

	"github.com/bigrocs/alipay/requests"
)

func TestMchPayMicropay(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AppId = "2016080900197401"
	client.Config.Method = "alipay.trade.pay"
	client.Config.PrivateKey = "MIIEogIBAAKCAQEAgc4wW8rlV9XvywSUJYsib7Csg6waonzGAISqZjDNqcDY2oYK/yTwHd9L1qH69fS6sTDY32R0oSFy9Pg22f040GAALECRTtcGkyBC1wsfQUICxhTASbOX1NzY4Hk//8+4dcNAGAVxeByLx52hKJyCMjDveTwZLRCijxsJ56kgcjulptA5LxsiImInxHyH5Ne+kgPcZUt4o5IbpcngpNZ2GVrASbPzV/5BiuBza883ey1c4Ps/dHjA7vNOjfJ/lpIxWLdmfrs1aHTkZdRsRcsvPWzBIYZ1mpUIlGF48S1hkptwA8BpNCR2Lhn1NO8fQkZkDxHxhBp5IFtp92bJ+PJ1iwIDAQABAoIBAAKhWvmBMwSEoUdL4HSiTSBWRCim5CNGw/xes9U+R+yYq8xByxdAeF0imTbQMXWqb94HT123e2WId/vey72/E9elGlvLMyaV1NxGWxRgdVGtNMEuynaz3O/DSeHCkzlrzCEWw8W41oPIkEyLNSjNgxXhR/j87sWrAKBFu/2lv7KQ23Me6rY5hiJq7RW+5z0deEqDLkPHmbM12lJoNAvmhX4Q715s1LivUsFP5KZkxKzGgh233S2BcGuTIYj036g1DI83GbG1cVSDW9gU+J43CKKOYu+dCqUpYBXvRoTFDySqTo6blv4iQD/rvDvqFJkNBsEHV2HNbLNK/fREpK43IwECgYEAwqRwjp5+osbYAtHSz+GEFoWgTozKSxYHmYc55xRvi9bE328Y4ZxyyXZiQCxFKwdmCQl/E2sm1MalRwKRRhDRnIgkIIS/wGuwJ7fWkeNjmO68MaExs/mmdn9eYQJA8knms/sB19qAFuPl9uSiY5CJfLxaTqqTzF1fM02ot3uM/7MCgYEAqrlvgXo0VPWNUxVdMRHXPaRk8kl+lFYb8qXJVU2X/vdShHJ4iPVeFUI/XTeOFWherHkH83u1BWhkLU+KvrYxyG0rxdINYI/priMKDjQ06tH2cQshbe1D66ku9TnM1kbiYN6xJpu3MJsre0v/hqgr+jQ9ucMCVkMLZRIb9i4ChskCgYBWJxLrDZxf0EOse9Mj0F322g6qtgaUVZvniMNIVPNKojVh8HrC4cys/4ldfjrfYNb5CQsGPVditspRNAG5UZh2AIx9GEUHlqLR8b03tb5P4tvJ2990GfxkVtwfdC/rDmrfHyshr8UiXJ1dOrXl/APfAj+2sinZOzr4KleTX0x5oQKBgFFEC8v6O8blS+xskvu0nlx9UH/0dAhwJMWQHRI56Uw4tlLNmoq3IZ3E9xyMQVn3YHmA80P3cuesFWNsJYM6fuAE930my69XUcjObQ7t0vKkF0cgIT2OX3JiCjQ420R0YXXzCyhxnvXIJx59wph6nNRw4aD3LrmZGGd4A09f/1vhAoGANc3hlyUk28qctgQQSi10/Ng5Fy4YIKW+g8WwNZ3voaeIyaSM/9jJ+teqqnJw8mcgGYo+YMqU+7IEwzrcjib4id2cu4MhlRswXKgaG2gSIUN7JEivjQisM6UUkXZtD91ZKx9JxTPVRYXqAjCcd6YPYt/C8PJTfGpQUFttX498P5g="
	client.Config.Sandbox = true
	// 配置参数
	request := requests.NewCommonRequest()
	request.BizContent = map[string]interface{}{
		"subject":         "测试商品名称1",
		"scene":           "bar_code",
		"auth_code":       "286906367716990981",
		"out_trade_no":    "GZ2020010117534314513",
		"total_amount":    0.02,
		"timeout_express": "30m",
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	// fmt.Println(response, err)
	fmt.Println(req, err)
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
