package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PaymentCallbacksSearch struct {
	request.PageInfo
	Id              int        `json:"id" form:"id"`                           // 主键ID
	MerchantOrderNo string     `json:"merchantOrderNo" form:"merchantOrderNo"` // 商户订单号
	OrderNo         string     `json:"orderNo" form:"orderNo"`                 // 平台订单号
	CallbackType    int        `json:"callbackType" form:"callbackType"`       // 回调类型
	MerchantId      string     `json:"merchantId" form:"merchantId"`           // 商户ID
	Amount          int64      `json:"amount" form:"amount"`                   // 金额
	Currency        string     `json:"currency" form:"currency"`               // 货币类型
	Status          string     `json:"status" form:"status"`                   // 状态
	PayType         string     `json:"payType" form:"payType"`                 // 支付方式
	RefCpf          string     `json:"refCpf" form:"refCpf"`                   // 参考CPF
	RefName         string     `json:"refName" form:"refName"`                 // 参考姓名
	ErrorMsg        string     `json:"errorMsg" form:"errorMsg"`               // 错误信息
	CallbackData    string     `json:"callbackData" form:"callbackData"`       // 回调数据
	Sign            string     `json:"sign" form:"sign"`                       // 签名
	SignVerified    bool       `json:"signVerified" form:"signVerified"`       // 签名验证
	IpAddress       string     `json:"ipAddress" form:"ipAddress"`             // IP地址
	UserAgent       string     `json:"userAgent" form:"userAgent"`             // 用户代理
	Processed       bool       `json:"processed" form:"processed"`             // 是否已处理
	ProcessedTime   *time.Time `json:"processedTime" form:"processedTime"`     // 处理时间
	RetryCount      int        `json:"retryCount" form:"retryCount"`           // 重试次数
	LastRetryTime   *time.Time `json:"lastRetryTime" form:"lastRetryTime"`     // 最后重试时间
	ErrorReason     string     `json:"errorReason" form:"errorReason"`         // 错误原因
	Remark          string     `json:"remark" form:"remark"`                   // 备注
	CreatedAtStart  *time.Time `json:"createdAtStart" form:"createdAtStart"`   // 创建时间开始
	CreatedAtEnd    *time.Time `json:"createdAtEnd" form:"createdAtEnd"`       // 创建时间结束
	UpdatedAtStart  *time.Time `json:"updatedAtStart" form:"updatedAtStart"`   // 更新时间开始
	UpdatedAtEnd    *time.Time `json:"updatedAtEnd" form:"updatedAtEnd"`       // 更新时间结束
}
