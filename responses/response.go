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
	WAITING    = "WAITING"    // 2	系统执行中请等待
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

// GetVerifySignDataMap 获取 GetVerifySignDataMap 校验后数据数据
func (res *CommonResponse) GetVerifySignDataMap() (m mxj.Map, err error) {
	r, err := res.GetHttpContentMap()
	if err != nil {
		return r, err
	}
	if r["sign"] != nil {
		ok, err := util.VerifySign(res.GetSignData(), r["sign"].(string), res.Config.AliPayPublicKey, res.Config.SignType)
		if err != nil {
			return r, err
		}
		if ok {
			return res.GetSignDataMap()
		}
	} else {
		return r, errors.New("sign is not")
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

// data{
// 	channel			//	通道内容		alipay、wechat、icbc
// 	content			//	第三方返回内容 	{}
// 	return_code		//	返回代码 		SUCCESS
// 	return_msg		//	返回消息		支付失败
// 	status			//	下单状态 		【SUCCESS成功、CLOSED关闭、USERPAYING等待用户付款、WAITING系统繁忙稍后查询】
// 	total_fee		//  订单金额		88
// 	refund_fee 		//  退款金额		10
// 	trade_no 		// 	渠道交易编号 	2013112011001004330000121536
// 	out_trade_no	// 	商户订单号		T1024501231476
//  out_refund_no	//  商户退款单号	T1024501231476_T
// 	wechat_open_id		//  微信openid		[oUpF8uN95-Ptaags6E_roPHg7AG
//  wechat_is_subscribe 	//  微信是否微信关注公众号
// 	alipay_logon_id  //	支付宝账号		158****1562
//  alipay_user_id  //	买家在支付宝的用户id	2088101117955611
// 	time_end		//  支付完成时间	20141030133525
// }
// GetSignDataMap 获取 MAP 数据
func (res *CommonResponse) GetSignDataMap() (mxj.Map, error) {
	data := mxj.New()
	content, err := mxj.NewMapJson([]byte(res.GetSignData()))
	if err != nil {
		return nil, err
	}
	if res.Request.ApiName == "alipay.trade.pay" {
		data = res.handerAlipayTradePay(content)
	}
	if res.Request.ApiName == "alipay.trade.query" {
		data = res.handerAlipayTradeQuery(content)
	}
	if res.Request.ApiName == "alipay.trade.refund" {
		data = res.handerAlipayTradeRefund(content)
	}
	if res.Request.ApiName == "alipay.trade.fastpay.refund.query" {
		data = res.handerAlipayTradeRefundQuery(content)
	}

	data["channel"] = "alipay" //渠道
	data["content"] = content
	return data, err
}

func (res *CommonResponse) handerAlipayTradePay(content mxj.Map) mxj.Map {
	data := mxj.New()
	data["status"] = "" // 状态
	if sub_msg, ok := content["sub_msg"]; ok {
		data["return_msg"] = sub_msg
	} else {
		data["return_msg"] = content["msg"]
	}
	if content["code"] == "10000" {
		data["return_code"] = SUCCESS
		data["status"] = SUCCESS
		data["total_fee"] = content["total_amount"]
		data["trade_no"] = content["trade_no"]
		data["out_trade_no"] = content["out_trade_no"]
		data["alipay_logon_id"] = content["buyer_logon_id"]
		data["alipay_user_id"] = content["buyer_user_id"]
		if t, ok := content["gmt_payment"]; ok {
			timeFormat := "2006-01-02 15:04:05"
			tt, _ := time.ParseInLocation(timeFormat, t.(string), time.Local)
			data["time_end"] = tt.Format("20060102150405")
		}
	} else {
		data["return_code"] = "FAIL"
	}
	return data
}

func (res *CommonResponse) handerAlipayTradeQuery(content mxj.Map) mxj.Map {
	data := mxj.New()
	data["status"] = "" // 状态
	if sub_msg, ok := content["sub_msg"]; ok {
		data["return_msg"] = sub_msg
	} else {
		data["return_msg"] = content["msg"]
	}
	if content["code"] == "10000" {
		data["return_code"] = SUCCESS
		switch content["trade_status"] {
		case "TRADE_CLOSED":
			data["status"] = CLOSED
		case "WAIT_BUYER_PAY":
			data["status"] = USERPAYING
		case "TRADE_SUCCESS":
			data["status"] = SUCCESS
		case "TRADE_FINISHED":
			data["status"] = SUCCESS
		}
		data["total_fee"] = content["total_amount"]
		data["trade_no"] = content["trade_no"]
		data["out_trade_no"] = content["out_trade_no"]
		data["alipay_logon_id"] = content["buyer_logon_id"]
		data["alipay_user_id"] = content["buyer_user_id"]
		if t, ok := content["send_pay_date"]; ok {
			timeFormat := "2006-01-02 15:04:05"
			tt, _ := time.ParseInLocation(timeFormat, t.(string), time.Local)
			data["time_end"] = tt.Format("20060102150405")
		}
	} else {
		data["return_code"] = "FAIL"
		if content["sub_code"] == "ACQ.TRADE_NOT_EXIST" {
			data["return_code"] = SUCCESS
			data["status"] = CLOSED
		}
	}
	return data
}

func (res *CommonResponse) handerAlipayTradeRefund(content mxj.Map) mxj.Map {
	data := mxj.New()
	if sub_msg, ok := content["sub_msg"]; ok {
		data["return_msg"] = sub_msg
	} else {
		data["return_msg"] = content["msg"]
	}
	if content["code"] == "10000" {
		data["return_code"] = SUCCESS
		data["refund_fee"] = content["refund_fee"]
		data["trade_no"] = content["trade_no"]
		data["out_trade_no"] = content["out_trade_no"]
		data["out_refund_no"] = content["out_request_no"]
	} else {
		data["return_code"] = "FAIL"
	}
	return data
}

func (res *CommonResponse) handerAlipayTradeRefundQuery(content mxj.Map) mxj.Map {
	data := mxj.New()
	if sub_msg, ok := content["sub_msg"]; ok {
		data["return_msg"] = sub_msg
	} else {
		data["return_msg"] = content["msg"]
	}
	if content["code"] == "10000" {
		if refund_amount, ok := content["refund_amount"]; ok {
			data["total_fee"] = content["total_amount"]
			data["refund_fee"] = refund_amount
			data["trade_no"] = content["trade_no"]
			data["out_trade_no"] = content["out_trade_no"]
			data["out_refund_no"] = content["out_request_no"]
			data["status"] = SUCCESS
		}
		data["return_code"] = SUCCESS
	} else {
		if content["sub_code"] == "ACQ.TRADE_NOT_EXIST" || content["sub_code"] == "TRADE_NOT_EXIST" {
			data["status"] = CLOSED
		}
		data["return_code"] = "FAIL"
	}
	return data
}
