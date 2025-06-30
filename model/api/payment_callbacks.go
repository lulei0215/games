// PaymentCallbacks
package api

import (
	"time"
)

// paymentCallbacks表   PaymentCallbacks
type PaymentCallbacks struct {
	Id              int        `json:"id" form:"id" gorm:"primarykey;comment:主键ID;column:id;size:20;"`                                                                                 //主键ID
	MerchantOrderNo string     `json:"merchantOrderNo" form:"merchantOrderNo" gorm:"comment:商户订单号;column:merchant_order_no;size:64;"`                                                  //商户订单号
	OrderNo         string     `json:"orderNo" form:"orderNo" gorm:"comment:平台订单号;column:order_no;size:64;"`                                                                           //平台订单号
	CallbackType    int        `json:"callbackType" form:"callbackType" gorm:"comment:回调类型: 1=充值回调 2=提现回调;column:callback_type;"`                                                      //回调类型: 1=充值回调 2=提现回调
	MerchantId      string     `json:"merchantId" form:"merchantId" gorm:"comment:商户ID;column:merchant_id;size:32;"`                                                                   //商户ID
	Amount          int64      `json:"amount" form:"amount" gorm:"comment:交易金额(分为单位);column:amount;size:19;"`                                                                          //交易金额(分为单位)
	Currency        string     `json:"currency" form:"currency" gorm:"comment:货币类型;column:currency;size:10;"`                                                                          //货币类型
	Status          string     `json:"status" form:"status" gorm:"comment:回调状态: PAID=支付成功 PAY_FAILED=支付失败 SUCCESS=提现成功 FAILED=提现失败 COMPLETED=已完成 REJECTED=已拒绝;column:status;size:20;"` //回调状态: PAID=支付成功 PAY_FAILED=支付失败 SUCCESS=提现成功 FAILED=提现失败 COMPLETED=已完成 REJECTED=已拒绝
	PayType         string     `json:"payType" form:"payType" gorm:"comment:支付方式(充值回调);column:pay_type;size:32;"`                                                                      //支付方式(充值回调)
	RefCpf          string     `json:"refCpf" form:"refCpf" gorm:"comment:参考CPF(充值回调);column:ref_cpf;size:32;"`                                                                        //参考CPF(充值回调)
	RefName         string     `json:"refName" form:"refName" gorm:"comment:参考姓名(充值回调);column:ref_name;size:128;"`                                                                     //参考姓名(充值回调)
	ErrorMsg        string     `json:"errorMsg" form:"errorMsg" gorm:"comment:错误信息(失败时);column:error_msg;size:500;"`                                                                   //错误信息(失败时)
	CallbackData    string     `json:"callbackData" form:"callbackData" gorm:"comment:回调原始数据(JSON格式);column:callback_data;type:json;"`                                                 //回调原始数据(JSON格式)
	Sign            string     `json:"sign" form:"sign" gorm:"comment:回调签名;column:sign;size:64;"`                                                                                      //回调签名
	SignVerified    bool       `json:"signVerified" form:"signVerified" gorm:"comment:签名验证结果: 0=未验证 1=验证成功 2=验证失败;column:sign_verified;"`                                              //签名验证结果: 0=未验证 1=验证成功 2=验证失败
	IpAddress       string     `json:"ipAddress" form:"ipAddress" gorm:"comment:回调来源IP;column:ip_address;size:45;"`                                                                    //回调来源IP
	UserAgent       string     `json:"userAgent" form:"userAgent" gorm:"comment:User-Agent;column:user_agent;size:500;"`                                                               //User-Agent
	Processed       bool       `json:"processed" form:"processed" gorm:"comment:处理状态: 0=未处理 1=已处理 2=处理失败;column:processed;"`                                                           //处理状态: 0=未处理 1=已处理 2=处理失败
	ProcessedTime   *time.Time `json:"processedTime" form:"processedTime" gorm:"comment:处理时间;column:processed_time;"`                                                                  //处理时间
	RetryCount      int        `json:"retryCount" form:"retryCount" gorm:"comment:重试次数;column:retry_count;size:10;"`                                                                   //重试次数
	LastRetryTime   *time.Time `json:"lastRetryTime" form:"lastRetryTime" gorm:"comment:最后重试时间;column:last_retry_time;"`                                                               //最后重试时间
	ErrorReason     string     `json:"errorReason" form:"errorReason" gorm:"comment:处理失败原因;column:error_reason;size:500;"`                                                             //处理失败原因
	Remark          string     `json:"remark" form:"remark" gorm:"comment:备注信息;column:remark;size:255;"`                                                                               //备注信息
	CreatedAt       *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`                                                                              //创建时间
	UpdatedAt       *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`                                                                              //更新时间
}

// TableName paymentCallbacks表 PaymentCallbacks payment_callbacks
func (PaymentCallbacks) TableName() string {
	return "payment_callbacks"
}
