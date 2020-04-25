package util

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"hash"
	"sort"
	"strconv"
	"strings"
)

const (
	SignType_MD5    = "MD5"
	SignType_SHA1   = "SHA1"
	SignType_SHA256 = "SHA256"
)

// VerifySign 验证支付
func VerifySign(params map[string]interface{}, privateKey string, signType string) bool {
	bodySign := params["sign"]
	sign, err := Sign(params, privateKey, signType)
	if err != nil {
		return false
	}
	return bodySign == sign
}

// EncodeSignParams 编码符号参数
func EncodeSignParams(params map[string]interface{}) string {
	var buf strings.Builder
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(InterfaceToString(v))
		buf.WriteByte('&')
	}
	return buf.String()[:buf.Len()-1]
}

// Sign 支付宝签名支付签名.
//  params: 待签名的参数集合
//  privateKey: api密钥
func Sign(params map[string]interface{}, privateKey string, signType string) (sign string, err error) {
	encodeSignParams := EncodeSignParams(params)
	var (
		block          *pem.Block
		h              hash.Hash
		key            *rsa.PrivateKey
		hashs          crypto.Hash
		encryptedBytes []byte
	)

	if block, _ = pem.Decode([]byte(privateKey)); block == nil {
		return "", errors.New("pem.Decode：privateKey decode error")
	}
	if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return
	}
	switch signType {
	case "RSA":
		h = sha1.New()
		hashs = crypto.SHA1
	case "RSA2":
		h = sha256.New()
		hashs = crypto.SHA256
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}
	if _, err = h.Write([]byte(encodeSignParams)); err != nil {
		return
	}
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, key, hashs, h.Sum(nil)); err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

// Sign2 支付宝签名支付签名.
//  params: 待签名的参数集合
//  privateKey: api密钥
//  h:      hash.Hash, 如果为 nil 则默认用 md5.New(), 特别注意 h 必须是 initial state.
func Sign2(params map[string]interface{}, privateKey string, h hash.Hash) string {
	if h == nil {
		h = md5.New()
	}

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	bufw := bufio.NewWriterSize(h, 128)
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		bufw.WriteString(k)
		bufw.WriteByte('=')
		bufw.WriteString(InterfaceToString(v))
		bufw.WriteByte('&')
	}
	bufw.WriteString("key=")
	bufw.WriteString(privateKey)
	bufw.Flush()

	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

// ParseNotifyResult 解析异步通知
func InterfaceToString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case float32:
		return strconv.FormatFloat(v.(float64), 'E', -1, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'E', -1, 64)
	}
	return ""
}

// jssdk 支付签名, signType 只支持 "MD5", "SHA1", 传入其他的值会 panic.
func JsapiSign(appId, timeStamp, nonceStr, packageStr, signType string, privateKey string) string {
	var h hash.Hash
	switch signType {
	case SignType_MD5:
		h = md5.New()
	case SignType_SHA1:
		h = sha1.New()
	default:
		panic("unsupported signType")
	}
	bufw := bufio.NewWriterSize(h, 128)

	// appId
	// nonceStr
	// package
	// signType
	// timeStamp
	bufw.WriteString("appId=")
	bufw.WriteString(appId)
	bufw.WriteString("&nonceStr=")
	bufw.WriteString(nonceStr)
	bufw.WriteString("&package=")
	bufw.WriteString(packageStr)
	bufw.WriteString("&signType=")
	bufw.WriteString(signType)
	bufw.WriteString("&timeStamp=")
	bufw.WriteString(timeStamp)
	bufw.WriteString("&key=")
	bufw.WriteString(privateKey)

	bufw.Flush()
	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

// EditAddressSign 收货地址共享接口签名
func EditAddressSign(appId, url, timestamp, nonceStr, accessToken string) string {
	h := sha1.New()
	bufw := bufio.NewWriterSize(h, 128)

	// accesstoken
	// appid
	// noncestr
	// timestamp
	// url
	bufw.WriteString("accesstoken=")
	bufw.WriteString(accessToken)
	bufw.WriteString("&appid=")
	bufw.WriteString(appId)
	bufw.WriteString("&noncestr=")
	bufw.WriteString(nonceStr)
	bufw.WriteString("&timestamp=")
	bufw.WriteString(timestamp)
	bufw.WriteString("&url=")
	bufw.WriteString(url)

	bufw.Flush()
	return hex.EncodeToString(h.Sum(nil))
}
