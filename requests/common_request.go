package requests

import (
	"errors"

	"github.com/bigrocs/alipay/util"
	"github.com/clbanning/mxj"
)

// CommonRequest 公共请求
type CommonRequest struct {
	BizContent map[string]interface{}
}

// NewCommonRequest 创建新的公共连接
func NewCommonRequest() (request *CommonRequest) {
	request = &CommonRequest{}
	return
}

// Request 公共请求和数据处理
func Request(request *CommonRequest, aliPayPublicKey string, signType string) (req mxj.Map, err error) {
	// 请求
	response, err := srv.Client.ProcessCommonRequest(request)
	if err != nil {
		return req, err
	}
	req, err = response.GetHttpContentMap()
	if ok, err := util.VerifySign(response.GetSignData(), req["sign"].(string), aliPayPublicKey, signType); !ok {
		if err != nil {
			return req, err
		}
		return req, errors.New("返回数据 Sign 校验失败")
	}
	return req, err
}
