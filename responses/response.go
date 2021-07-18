/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package responses

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/clbanning/mxj"

	"github.com/bigrocs/alipay/config"
	"github.com/bigrocs/alipay/requests"
	"github.com/bigrocs/alipay/util"
)

const (
	CLOSED     = "CLOSED"     // -1 订单关闭
	USERPAYING = "USERPAYING" // 0	订单支付中
	SUCCESS    = "SUCCESS"    // 1	订单支付成功
	WAIT       = "WAIT"       // 2	系统执行中请等待
)

// CommonResponse 公共回应
type CommonResponse struct {
	Config      *config.Config
	Request     *requests.CommonRequest
	httpContent []byte
	json        string
}

type Map *mxj.Map

// NewCommonResponse 创建新的请求返回
func NewCommonResponse(config *config.Config, request *requests.CommonRequest) (response *CommonResponse) {
	c := &CommonResponse{}
	c.Config = config
	c.Request = request
	return c
}

// GetHttpContentJson 获取 JSON 数据
func (res *CommonResponse) GetHttpContentJson() string {
	return res.json
}

// GetHttpContentMap 获取 MAP 数据
func (res *CommonResponse) GetHttpContentMap() (mxj.Map, error) {
	return mxj.NewMapJson([]byte(res.json))
}

// data{
// 	channel			//	通道内容		alipay、wechat、icbc
// 	content			//	第三方返回内容 	{}
// 	return_code		//	返回代码 		SUCCESS
// 	return_msg		//	返回消息		支付失败
// 	stauts			//	下单状态 		【SUCCESS成功、CLOSED关闭、USERPAYING等待用户付款、WAIT系统繁忙稍后查询】
// 	total_fee		//  订单金额		88
// 	refund_fee 		//  退款金额		10
// 	trade_no 		// 	渠道交易编号 	2013112011001004330000121536
// 	out_trade_no	// 	商户订单号		T1024501231476
// 	sub_open_id		//  微信openid		[oUpF8uN95-Ptaags6E_roPHg7AG
// 	buyer_logon_id  //	支付宝账号		158****1562
//  buyer_user_id  //	买家在支付宝的用户id	2088101117955611
// 	time_end		//  支付完成时间	20141030133525
// }

// GetSignDataMap 获取 MAP 数据
func (res *CommonResponse) GetSignDataMap() (mxj.Map, error) {
	data := mxj.New()
	content, err := mxj.NewMapJson([]byte(res.GetSignData()))
	data["channel"] = "alipay" //渠道
	data["content"] = content
	data["return_code"] = "" // 返回代码
	data["return_msg"] = ""  // 返回消息
	data["stauts"] = ""      // 状态

	// 下单
	// 查询 交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
	if sub_msg, ok := content["sub_msg"]; ok {
		data["return_msg"] = sub_msg
	} else {
		data["return_msg"] = content["msg"]
	}
	if content["code"] == "10000" {
		data["return_code"] = SUCCESS
		if res.Request.ApiName == "alipay.trade.pay" {
			data["stauts"] = SUCCESS
		}
		switch content["trade_status"] {
		case "TRADE_CLOSED":
			data["stauts"] = CLOSED
		case "WAIT_BUYER_PAY":
			data["stauts"] = USERPAYING
		case "TRADE_SUCCESS":
			data["stauts"] = SUCCESS
		case "TRADE_FINISHED":
			data["stauts"] = SUCCESS
		}
	} else {
		data["return_code"] = "FAIL"
		// 系列关闭订单属于正常
		if content["code"] == "10003" { // 下单等待用户付款
			data["return_code"] = SUCCESS
			data["stauts"] = USERPAYING
		}
		if content["sub_code"] == "ACQ.TRADE_HAS_SUCCESS" {
			data["return_code"] = SUCCESS
			data["stauts"] = SUCCESS
		}
		if content["sub_code"] == "ACQ.TRADE_HAS_CLOSE" {
			data["return_code"] = SUCCESS
			data["stauts"] = CLOSED
		}
		if content["sub_code"] == "ACQ.TRADE_NOT_EXIST" {
			data["return_code"] = SUCCESS
			data["stauts"] = CLOSED
		}
	}
	data["total_fee"] = content["total_amount"]
	data["refund_fee"] = content["refund_fee"]
	if f, ok := content["refund_amount"]; ok {
		data["refund_fee"] = f
	}
	data["trade_no"] = content["trade_no"]
	data["out_trade_no"] = content["out_trade_no"]
	data["buyer_logon_id"] = content["buyer_logon_id"]
	data["buyer_user_id"] = content["buyer_user_id"]
	if t, ok := content["gmt_payment"]; ok {
		timeFormat := "2006-01-02 15:04:05"
		tt, _ := time.ParseInLocation(timeFormat, t.(string), time.Local)
		data["time_end"] = tt.Format("20060102150405")
	}
	return data, err
}

// GetVerifySignDataMap 获取 GetVerifySignDataMap 校验后数据数据
func (res *CommonResponse) GetVerifySignDataMap() (m mxj.Map, err error) {
	r, err := res.GetHttpContentMap()
	if err != nil {
		return m, err
	}
	if r["sign"] != nil {
		ok, err := util.VerifySign(res.GetSignData(), r["sign"].(string), res.Config.AliPayPublicKey, res.Config.SignType)
		if err != nil {
			return m, err
		}
		if ok {
			return res.GetSignDataMap()
		}
	} else {
		return m, errors.New("sign is not")
	}
	return
}

// GetSignData 获取 SignData 数据
func (res *CommonResponse) GetSignData() string {
	return util.GetSignData(res.json)
}

// GetSign 获取 Sign 数据
func (res *CommonResponse) GetSign() (string, error) {
	mv, err := res.GetHttpContentMap()
	if err != nil {
		return "", err
	}
	if _, ok := mv["sign"]; ok { //去掉 xml 外层
		return mv["sign"].(string), err
	}
	return "", err
}

// SetHttpContent 设置请求信息
func (res *CommonResponse) SetHttpContent(httpContent []byte, dataType string) {
	res.httpContent = httpContent
	switch dataType {
	case "xml":
		mv, _ := mxj.NewMapXml(res.httpContent) // unmarshal
		var str interface{}
		if _, ok := mv["xml"]; ok { //去掉 xml 外层
			str = mv["xml"]
		} else {
			str = mv
		}
		jsonStr, _ := json.Marshal(str)
		res.json = string(jsonStr)
	case "string":
		res.json = string(res.httpContent)
	}
}
