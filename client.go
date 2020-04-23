package alipay

import (
	"github.com/bigrocs/alipay/common"
	"github.com/bigrocs/alipay/config"
	"github.com/bigrocs/alipay/requests"
	"github.com/bigrocs/alipay/responses"
)

// Client the type Client
type Client struct {
	Config *config.Config
}

// NewClient 创建默认连接
func NewClient() (client *Client) {
	client = &Client{
		Config: &config.Config{},
	}
	return
}

// ProcessCommonRequest 处理公共请求
func (client *Client) ProcessCommonRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	response = responses.NewCommonResponse()
	err = client.DoAction(request, response)
	return
}

// DoAction 执行动作
func (client *Client) DoAction(request *requests.CommonRequest, response *responses.CommonResponse) (err error) {
	// 创建访问链接
	u := &common.Common{
		Config:   client.Config,
		Requests: request,
	}
	err = u.Action(response)
	if err != nil {
		return err
	}
	return
}
