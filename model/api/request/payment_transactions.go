package request

import (
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PaymentTransactionsSearch struct {
	request.PageInfo
	Id              uint       `json:"id" form:"id"`                           // 主键ID
	UserId          uint       `json:"userId" form:"userId"`                   // 用户ID
	MerchantOrderNo string     `json:"merchantOrderNo" form:"merchantOrderNo"` // 商户订单号
	OrderNo         string     `json:"orderNo" form:"orderNo"`                 // 平台订单号
	TransactionType int        `json:"transactionType" form:"transactionType"` // 交易类型
	Amount          int        `json:"amount" form:"amount"`                   // 交易金额
	Currency        string     `json:"currency" form:"currency"`               // 货币类型
	Status          string     `json:"status" form:"status"`                   // 交易状态
	PayType         string     `json:"payType" form:"payType"`                 // 支付方式
	AccountType     string     `json:"accountType" form:"accountType"`         // 账户类型
	AccountNo       string     `json:"accountNo" form:"accountNo"`             // 账号
	AccountName     string     `json:"accountName" form:"accountName"`         // 账户名
	Content         string     `json:"content" form:"content"`                 // 交易内容/备注
	ClientIp        string     `json:"clientIp" form:"clientIp"`               // 客户端IP地址
	CallbackUrl     string     `json:"callbackUrl" form:"callbackUrl"`         // 回调地址
	RedirectUrl     string     `json:"redirectUrl" form:"redirectUrl"`         // 跳转地址
	PayUrl          string     `json:"payUrl" form:"payUrl"`                   // 支付链接
	PayRaw          string     `json:"payRaw" form:"payRaw"`                   // 支付原始数据
	ErrorMsg        string     `json:"errorMsg" form:"errorMsg"`               // 错误信息
	RefCpf          string     `json:"refCpf" form:"refCpf"`                   // 参考CPF
	RefName         string     `json:"refName" form:"refName"`                 // 参考姓名
	CreatedAtStart  *time.Time `json:"createdAtStart" form:"createdAtStart"`   // 创建时间开始
	CreatedAtEnd    *time.Time `json:"createdAtEnd" form:"createdAtEnd"`       // 创建时间结束
	UpdatedAtStart  *time.Time `json:"updatedAtStart" form:"updatedAtStart"`   // 更新时间开始
	UpdatedAtEnd    *time.Time `json:"updatedAtEnd" form:"updatedAtEnd"`       // 更新时间结束
	DeletedAtStart  *time.Time `json:"deletedAtStart" form:"deletedAtStart"`   // 删除时间开始
	DeletedAtEnd    *time.Time `json:"deletedAtEnd" form:"deletedAtEnd"`       // 删除时间结束
}
type CreateTradeData struct {
	Amount   int64  `json:"amount"`
	PayType  string `json:"pay_type"`
	Currency string `json:"currency"`
	Content  string `json:"content"`
	ClientIp string `json:"clientIp"`
	Callback string `json:"callback"`
	Redirect string `json:"redirect"`
}
type CreateTradeData2 struct {
	TotalAmount string `json:"totalAmount"`
	PayCardNo   string `json:"payCardNo"`
	PayBankCode string `json:"payBankCode"`
	PayName     string `json:"payName"`
	PayEmail    string `json:"payEmail"`
	PayPhone    string `json:"payPhone"`
	PayViewUrl  string `json:"payViewUrl"`
}
type CreatePaymentData struct {
	Amount        int64  `json:"amount"`
	AccountId     int64  `json:"accountId"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	AccountType   string `json:"accountType"`
	CpfNumber     string `json:"cpfNumber"`
	PaymentType   string `json:"paymentType"`
}
type CreatePaymentData2 struct {
	TotalAmount string `json:"totalAmount"`
	AccountId   string `json:"accountId"`
}

// 创建交易响应结构体
type CreateTradeResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		PayUrl          string      `json:"payUrl"`
		PayRaw          string      `json:"payRaw,omitempty"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		PayType         string      `json:"payType"`
	} `json:"data"`
}

// 查询交易订单响应结构体
type QueryTradeResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		PayType         string      `json:"payType"`
		RefCpf          string      `json:"ref_cpf,omitempty"`
		RefName         string      `json:"ref_name,omitempty"`
	} `json:"data"`
}

// 交易回调通知结构体 (JSON格式)
type TradeCallbackRequest struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		PayType         string      `json:"payType"`
		RefCpf          string      `json:"ref_cpf,omitempty"`
		RefName         string      `json:"ref_name,omitempty"`
		Sign            string      `json:"sign"`
	} `json:"data"`
}

// 交易回调通知结构体 (Form格式)
type TradeCallbackFormRequest struct {
	MerchantId      string `form:"merchantId"`
	MerchantOrderNo string `form:"merchantOrderNo"`
	OrderNo         string `form:"orderNo"`
	Amount          int64  `form:"amount"`
	Status          string `form:"status"`
	Currency        string `form:"currency"`
	PayType         string `form:"payType"`
	RefCpf          string `form:"ref_cpf"`
	RefName         string `form:"ref_name"`
	Sign            string `form:"sign"`
}

type TradeCallback2Request struct {
	MerNo        string `json:"merNo" binding:"required"`
	CurrencyCode string `json:"currencyCode" binding:"required"`
	PayType      string `json:"payType" binding:"required"`
	OutTradeNo   string `json:"outTradeNo" binding:"required"`
	TotalAmount  string `json:"totalAmount" binding:"required"`
	PayOrderNo   string `json:"payOrderNo" binding:"required"`
	PayState     string `json:"payState" binding:"required"`
	PayDate      string `json:"payDate" binding:"required"`
	Sign         string `json:"sign" binding:"required"`
}

type PaymentCallbackFormRequest struct {
	MerchantId      string `form:"merchantId"`
	MerchantOrderNo string `form:"merchantOrderNo"`
	OrderNo         string `form:"orderNo"`
	Amount          int64  `form:"amount"`
	Status          string `form:"status"`
	Currency        string `form:"currency"`
	ErrorMsg        string `form:"errorMsg"`
	Sign            string `form:"sign"`
}
type PaymentCallback2FormRequest struct {
	MerNo        string `json:"merNo" binding:"required"`
	CurrencyCode string `json:"currencyCode" binding:"required"`
	OutTradeNo   string `json:"outTradeNo" binding:"required"`
	TotalAmount  string `json:"totalAmount" binding:"required"`
	RemitOrderNo string `json:"remitOrderNo" binding:"required"`
	RemitState   string `json:"remitState" binding:"required"`
	RemitDate    string `json:"remitDate" binding:"required"`
	OrderMessage string `json:"orderMessage" binding:"required"`
	Sign         string `json:"sign" binding:"required"`
}

// 创建提现响应结构体
type CreatePaymentResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
	} `json:"data"`
}

// 查询提现响应结构体
type QueryPaymentResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
	} `json:"data"`
}

// 提现回调通知结构体
type PaymentCallbackRequest struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
		Sign            string      `json:"sign"`
	} `json:"data"`
}

// 余额查询响应结构体
type BalanceResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		Balance          int64  `json:"balance"`
		UnsettledBalance int64  `json:"unsettledBalance"`
		FrozenAmount     int64  `json:"frozenAmount"`
		Currency         string `json:"currency"`
	} `json:"data"`
}

// 提现反查响应结构体
type ReversePaymentResponse struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Data      struct {
		MerchantId      interface{} `json:"merchantId"`
		MerchantOrderNo string      `json:"merchantOrderNo"`
		OrderNo         string      `json:"orderNo"`
		Amount          int64       `json:"amount"`
		Status          string      `json:"status"`
		Currency        string      `json:"currency"`
		ErrorMsg        string      `json:"errorMsg,omitempty"`
	} `json:"data"`
}

type PaymentClient struct {
	BaseURL    string
	MerchantId string
	SecretKey  string
}

type CallbackDeduplicator struct {
	processedCallbacks map[string]bool
	mutex              sync.RWMutex
}
