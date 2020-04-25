# alipay SDK
### 微服务内核访问演示

#### 支付付款码支付
```
// 详情请查看 client_text.go 文件
{
		"subject":         "测试商品名称1",
		"scene":           "bar_code",
		"auth_code":       "286906367716990981",
		"out_trade_no":    "GZ2020010117534314513",
		"total_amount":    0.02,
		"timeout_express": "30m",
}
```
#### 支付付款码支付查询
```
// 详情请查看 client_text.go 文件
{
    "out_trade_no":     "202002100007",
}
```
- 具体参数参考支付开发文档
- https://api.wechat.com/
