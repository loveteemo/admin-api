package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type EmptyData interface{}

type Result struct {
	Code    int       `json:"code"`
	Data    EmptyData `json:"data"`
	Message string    `json:"message"`
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
