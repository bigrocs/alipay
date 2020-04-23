package common

import (
	"github.com/bigrocs/alipay/config"
	"github.com/bigrocs/alipay/requests"
	"github.com/bigrocs/alipay/responses"
)

// Common 公共封装
type Common struct {
	Config   *config.Config
	Requests *requests.CommonRequest
}

// Action 创建新的公共连接
func (c *Common) Action(response *responses.CommonResponse) (err error) {
	req := c.Requests
	// 根据作用域分发
	switch req.Domain {
	case "trade": // 支付接口
		m := &Trade{c}
		err = m.Request(response)
	}
	return
}

// 默认 API
func (c *Common) APIBaseURL() string { // TODO(): 后期做容灾功能
	return "https://openapi.alipay.com/gateway.do"
}
