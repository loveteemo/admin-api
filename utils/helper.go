package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type EmptyData interface{}

type Result struct {
	Code    int       `json:"code"`
	Data    EmptyData `json:"data"`
	Message string    `json:"message"`
}

type IpAddress struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

// GetAddress IP地址转成省市
func GetAddress(ip string) string {
	url := fmt.Sprintf("https://restapi.amap.com/v3/ip?key=64f192da1cb2edfddb573f5afa9097a6&ip=%s", ip)
	response, err := GetRequest(url)
	if err != nil {
		return "未知地址"
	}
	data := IpAddress{}
	if err := json.Unmarshal([]byte(response), &data); err != nil {
		return "未知地址"
	}

	return data.Province + data.City
}

// GetRequest GET请求
func GetRequest(url string) (data string, err error) {
	client := &http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	}
	return
}

// GenerateToken 生成 Token
func GenerateToken() string {
	// 获取当前时间和唯一 ID
	now := time.Now().UnixNano()
	uid := fmt.Sprintf("%d", now)

	// 将唯一 ID 和当前时间拼接在一起
	s := uid + fmt.Sprintf("%d", time.Now().Unix())

	// 对字符串进行 MD5 加密，并使用 Base64 编码
	h := md5.New()
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// EncryptPassword 加密密码
func EncryptPassword(password string, salt string) string {
	h := md5.New()
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}

// ReturnResult 返回结果
func ReturnResult(context *gin.Context, code int, message string, data interface{}) {

	if data == nil {
		data = make(map[string]interface{})
	}

	context.JSON(200, Result{
		Code:    code,
		Data:    data,
		Message: message,
	})
}
