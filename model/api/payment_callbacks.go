// PaymentCallbacks
package api

import (
	"time"
)

// paymentCallbacks表   PaymentCallbacks
type PaymentCallbacks struct {
	Id              int        `json:"id" form:"id" gorm:"primarykey;comment:id;column:id;size:20;"`                                            //id
	MerchantOrderNo string     `json:"merchantOrderNo" form:"merchantOrderNo" gorm:"comment:merchantOrderNo;column:merchant_order_no;size:64;"` //merchantOrderNo
	OrderNo         string     `json:"orderNo" form:"orderNo" gorm:"comment:orderNo;column:order_no;size:64;"`                                  //orderNo
	CallbackType    int        `json:"callbackType" form:"callbackType" gorm:"comment:callbackType;column:callback_type;"`                      //callbackType
	MerchantId      string     `json:"merchantId" form:"merchantId" gorm:"comment:merchantId;column:merchant_id;size:32;"`                      //merchantId
	Amount          int64      `json:"amount" form:"amount" gorm:"comment:amount;column:amount;size:19;"`                                       //amount
	Currency        string     `json:"currency" form:"currency" gorm:"comment:currency;column:currency;size:10;"`                               //currency
	Status          string     `json:"status" form:"status" gorm:"comment:status;column:status;size:20;"`                                       //status
	PayType         string     `json:"payType" form:"payType" gorm:"comment:payType;column:pay_type;size:32;"`                                  //payType
	RefCpf          string     `json:"refCpf" form:"refCpf" gorm:"comment:refCpf;column:ref_cpf;size:32;"`                                      //refCpf
	RefName         string     `json:"refName" form:"refName" gorm:"comment:refName;column:ref_name;size:128;"`                                 //refName
	ErrorMsg        string     `json:"errorMsg" form:"errorMsg" gorm:"comment:errorMsg;column:error_msg;size:500;"`                             //errorMsg
	CallbackData    string     `json:"callbackData" form:"callbackData" gorm:"comment:callbackData;column:callback_data;type:json;"`            //callbackData
	Sign            string     `json:"sign" form:"sign" gorm:"comment:sign;column:sign;size:64;"`                                               //sign
	SignVerified    bool       `json:"signVerified" form:"signVerified" gorm:"comment:signVerified;column:sign_verified;"`                      //signVerified
	IpAddress       string     `json:"ipAddress" form:"ipAddress" gorm:"comment:ipAddress;column:ip_address;size:45;"`                          //ipAddress
	UserAgent       string     `json:"userAgent" form:"userAgent" gorm:"comment:userAgent;column:user_agent;size:500;"`                         //userAgent
	Processed       bool       `json:"processed" form:"processed" gorm:"comment:processed;column:processed;"`                                   //processed
	ProcessedTime   *time.Time `json:"processedTime" form:"processedTime" gorm:"comment:processedTime;column:processed_time;"`                  //processedTime
	RetryCount      int        `json:"retryCount" form:"retryCount" gorm:"comment:retryCount;column:retry_count;size:10;"`                      //retryCount
	LastRetryTime   *time.Time `json:"lastRetryTime" form:"lastRetryTime" gorm:"comment:lastRetryTime;column:last_retry_time;"`                 //lastRetryTime
	ErrorReason     string     `json:"errorReason" form:"errorReason" gorm:"comment:errorReason;column:error_reason;size:500;"`                 //errorReason
	Remark          string     `json:"remark" form:"remark" gorm:"comment:remark;column:remark;size:255;"`                                      //remark
	CreatedAt       *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:createdAt;column:created_at;"`                                  //createdAt
	UpdatedAt       *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:updatedAt;column:updated_at;"`                                  //updatedAt
}

// TableName paymentCallbacks表 PaymentCallbacks payment_callbacks
func (PaymentCallbacks) TableName() string {
	return "payment_callbacks"
}
