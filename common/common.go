package common

import (
	"time"

	"github.com/clbanning/mxj"

	"github.com/bigrocs/alipay/config"
	"github.com/bigrocs/alipay/requests"
	"github.com/bigrocs/alipay/responses"
	"github.com/bigrocs/alipay/util"
)

// Common 公共封装
type Common struct {
	Config   *config.Config
	Requests *requests.CommonRequest
}

// Action 创建新的公共连接
func (c *Common) Action(response *responses.CommonResponse) (err error) {
	return c.Request(response)
}

// APIBaseURL 默认 API 网关
func (c *Common) APIBaseURL() string { // TODO(): 后期做容灾功能
	con := c.Config
	if con.Sandbox { // 沙盒模式
		return "https://openapi.alipaydev.com/gateway.do?charset=utf-8"
	}
	return "https://openapi.alipay.com/gateway.do?charset=utf-8"
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
func (c *Common) Request(response *responses.CommonResponse) (err error) {
	con := c.Config
	req := c.Requests
	bizContent, err := mxj.Map(req.BizContent).Json()
	if err != nil {
		return err
	}
	// 构建配置参数
	params := map[string]interface{}{
		"app_id":         con.AppId,
		"method":         req.ApiName,
		"sign_type":      con.SignType,
		"app_auth_token": con.AppAuthToken,
		"notify_url":     con.NotifyUrl,
		"format":         con.Format,
		"timestamp":      time.Now().Format("2006-01-02 15:04:05"),
		"biz_content":    string(bizContent),
	}
	if con.Format != "" {
		params["format"] = con.Format
	} else {
		params["format"] = "JSON"
	}
	if con.Charset != "" {
		params["charset"] = con.Charset
	} else {
		params["charset"] = "utf-8"
	}
	if con.SignType != "" {
		params["sign_type"] = con.SignType
	} else {
		params["sign_type"] = "RSA2"
	}
	if con.Version != "" {
		params["version"] = con.Version
	} else {
		params["version"] = "1.0"
	}
	sign, err := util.Sign(params, con.PrivateKey, params["sign_type"].(string)) // 开发签名
	if err != nil {
		return err
	}
	params["sign"] = sign
	urlParam := util.FormatURLParam(params)
	res, err := util.PostForm(c.APIBaseURL(), urlParam)
	if err != nil {
		return err
	}
	response.SetHttpContent(res, "string")
	return
}
