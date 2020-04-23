// 微信商户平台(微信支付) SDK
package common

import (
	"time"

	"github.com/bigrocs/alipay/responses"
	"github.com/bigrocs/alipay/util"
)

var apiUrlsTrade = map[string]string{
	"pay.micropay":   "/pay/micropay",   //付款码支付
	"pay.orderquery": "/pay/orderquery", //付款码支付查询
}

// Trade 支付结构
type Trade struct {
	c *Common
}

// Request 执行请求
// AppId        string `json:"app_id"`         //支付宝分配给开发者的应用ID
// Method       string `json:"method"`         //接口名称
// Format       string `json:"format"`         //仅支持 JSON
// Charset      string `json:"charset"`        //请求使用的编码格式，如utf-8,gbk,gb2312等，推荐使用 utf-8
// SignType     string `json:"sign_type"`      //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
// Sign         string `json:"sign"`           //商户请求参数的签名串
// Timestamp    string `json:"timestamp"`      //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
// Version      string `json:"version"`        //调用的接口版本，固定为：1.0
// NotifyUrl    string `json:"notify_url"`     //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
// BizContent   string `json:"biz_content"`    //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
// AppAuthToken string `json:"app_auth_token"` //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
// ReturnUrl    string `json:"return_url"`     //HTTP/HTTPS开头字符串
func (m *Trade) Request(response *responses.CommonResponse) (err error) {
	c := m.c.Config
	req := m.c.Requests
	// 构建配置参数
	req.QueryParams["app_id"] = c.AppId
	req.QueryParams["method"] = c.Method
	req.QueryParams["sign_type"] = c.SignType
	req.QueryParams["app_auth_token"] = c.AppAuthToken
	req.QueryParams["notify_url"] = c.NotifyUrl
	req.QueryParams["format"] = c.Format
	if c.Charset != "" {
		req.QueryParams["charset"] = c.Charset
	} else {
		req.QueryParams["charset"] = "utf-8"
	}
	if c.SignType != "" {
		req.QueryParams["sign_type"] = c.SignType
	} else {
		req.QueryParams["sign_type"] = "RSA2"
	}
	req.QueryParams["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	if c.Version != "" {
		req.QueryParams["version"] = c.Version
	} else {
		req.QueryParams["version"] = "1.0"
	}

	// req.QueryParams["sign"] = c.Sign // 开发签名
	if err != nil {
		return err
	}
	res, err := util.PostXML(m.c.APIBaseURL(), req.QueryParams)
	if err != nil {
		return err
	}
	response.SetHttpContent(res, "xml")
	return
}
