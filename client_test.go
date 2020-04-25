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

	// 配置参数
	request := requests.NewCommonRequest()
	request.BizContent = map[string]interface{}{
		"auth_code":    "285653237303565644",
		"title":        "测试商品名称1",
		"out_trade_no": "GZ202001011753431451",
		"total_amount": 0.01,
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	req, err := response.GetHttpContentMap()
	// fmt.Println(response, err)
	fmt.Println(req, err)
}
