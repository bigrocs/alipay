package config

type Config struct {
	AppId           string `json:"app_id"`             //支付宝分配给开发者的应用ID
	PrivateKey      string `json:"private_key"`        //私钥
	AliPayPublicKey string `json:"ali_pay_public_key"` //支付宝弓腰
	Method          string `json:"method"`             //接口名称
	Format          string `json:"format"`             //仅支持 JSON
	Charset         string `json:"charset"`            //请求使用的编码格式，如utf-8,gbk,gb2312等，推荐使用 utf-8
	SignType        string `json:"sign_type"`          //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
	Sign            string `json:"sign"`               //商户请求参数的签名串
	Timestamp       string `json:"timestamp"`          //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version         string `json:"version"`            //调用的接口版本，固定为：1.0
	NotifyUrl       string `json:"notify_url"`         //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	BizContent      string `json:"biz_content"`        //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
	AppAuthToken    string `json:"app_auth_token"`     //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}
