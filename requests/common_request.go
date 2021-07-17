package requests

// CommonRequest 公共请求
type CommonRequest struct {
	Domain     string
	ApiName    string
	BizContent map[string]interface{}
}

// NewCommonRequest 创建新的公共连接
func NewCommonRequest() (request *CommonRequest) {
	request = &CommonRequest{}
	return
}
