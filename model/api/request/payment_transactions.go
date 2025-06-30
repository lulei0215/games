package request

import (
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PaymentTransactionsSearch struct {
	request.PageInfo
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
type CreatePaymentData struct {
	Amount int64  `json:"amount"`
	Id     string `json:"id"`
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

// 交易回调通知结构体
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
