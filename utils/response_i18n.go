package utils

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/i18n"
	"github.com/gin-gonic/gin"
)

// ResponseI18n 多语言响应结构
type ResponseI18n struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// GetClientLang 从请求头获取客户端语言
func GetClientLang(c *gin.Context) string {
	// 优先从自定义头获取
	if lang := c.GetHeader("X-Language"); lang != "" {
		return i18n.NormalizeLang(lang)
	}

	// 从 Accept-Language 头获取
	if acceptLang := c.GetHeader("Accept-Language"); acceptLang != "" {
		return i18n.GetLangFromHeader(acceptLang)
	}

	return i18n.DefaultLang
}

// ResultI18n 多语言响应
func ResultI18n(code int, data interface{}, msgKey i18n.MessageKey, c *gin.Context) {
	lang := GetClientLang(c)
	message := i18n.GetMessage(lang, msgKey)

	c.JSON(http.StatusOK, ResponseI18n{
		Code: code,
		Data: data,
		Msg:  message,
	})
}

// OkI18n 成功响应
func OkI18n(c *gin.Context) {
	ResultI18n(0, map[string]interface{}{}, i18n.MsgSuccess, c)
}

// OkWithDataI18n 成功响应带数据
func OkWithDataI18n(data interface{}, c *gin.Context) {
	ResultI18n(0, data, i18n.MsgSuccess, c)
}

// OkWithMessageI18n 成功响应带自定义消息
func OkWithMessageI18n(msgKey i18n.MessageKey, c *gin.Context) {
	ResultI18n(0, map[string]interface{}{}, msgKey, c)
}

// OkWithDetailedI18n 成功响应带数据和自定义消息
func OkWithDetailedI18n(data interface{}, msgKey i18n.MessageKey, c *gin.Context) {
	ResultI18n(0, data, msgKey, c)
}

// FailI18n 失败响应
func FailI18n(c *gin.Context) {
	ResultI18n(400, map[string]interface{}{}, i18n.MsgFailed, c)
}

// FailWithMessageI18n 失败响应带消息
func FailWithMessageI18n(msgKey i18n.MessageKey, c *gin.Context) {
	ResultI18n(400, map[string]interface{}{}, msgKey, c)
}

// FailWithDetailedI18n 失败响应带数据和消息
func FailWithDetailedI18n(data interface{}, msgKey i18n.MessageKey, c *gin.Context) {
	ResultI18n(400, data, msgKey, c)
}

// UnauthorizedI18n 未授权响应
func UnauthorizedI18n(c *gin.Context) {
	ResultI18n(401, map[string]interface{}{}, i18n.MsgUserNotFound, c)
}

// ServerErrorI18n 服务器错误响应
func ServerErrorI18n(msgKey i18n.MessageKey, c *gin.Context) {
	ResultI18n(500, map[string]interface{}{}, msgKey, c)
}

// CustomErrorI18n 自定义错误码响应
func CustomErrorI18n(code int, msgKey i18n.MessageKey, c *gin.Context) {
	ResultI18n(code, map[string]interface{}{}, msgKey, c)
}
