// PaymentTransactions
package api

import (
	"time"
)

// paymentTransactions表   PaymentTransactions
type PaymentTransactions struct {
	Id              uint       `json:"id" form:"id" gorm:"primarykey;comment:主键ID;column:id;size:20;"`                                                                                                                          //主键ID
	UserId          uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:20;"`                                                                                                                        //用户ID
	MerchantOrderNo string     `json:"merchantOrderNo" form:"merchantOrderNo" gorm:"comment:商户订单号;column:merchant_order_no;size:64;"`                                                                                           //商户订单号
	OrderNo         string     `json:"orderNo" form:"orderNo" gorm:"comment:平台订单号;column:order_no;size:64;"`                                                                                                                    //平台订单号
	TransactionType int        `json:"transactionType" form:"transactionType" gorm:"comment:交易类型: 1=充值 2=提现;column:transaction_type;"`                                                                                          //交易类型: 1=充值 2=提现
	Amount          int        `json:"amount" form:"amount" gorm:"comment:交易金额(分为单位);column:amount;size:19;"`                                                                                                                   //交易金额(分为单位)
	Currency        string     `json:"currency" form:"currency" gorm:"comment:货币类型: BRL=雷亚尔 USD=美元 CNY=人民币 PKR=巴基斯坦卢比 PHP=菲律宾比索;column:currency;size:10;"`                                                                      //货币类型: BRL=雷亚尔 USD=美元 CNY=人民币 PKR=巴基斯坦卢比 PHP=菲律宾比索
	Status          string     `json:"status" form:"status" gorm:"comment:交易状态: WAITING_PAY=待支付 PAYING=支付中 PAID=支付成功 PAY_FAILED=支付失败 SUCCESS=提现成功 FAILED=提现失败 COMPLETED=已完成 REJECTED=已拒绝;column:status;size:20;"`               //交易状态: WAITING_PAY=待支付 PAYING=支付中 PAID=支付成功 PAY_FAILED=支付失败 SUCCESS=提现成功 FAILED=提现失败 COMPLETED=已完成 REJECTED=已拒绝
	PayType         string     `json:"payType" form:"payType" gorm:"comment:支付方式: PIX_QRCODE=巴西PIX扫码 PKR_EASYPAISA=EasyPaisa钱包 PKR_JAZZCASH=JazzCash钱包 PHQR=菲律宾扫码 GCASH=GCASH;column:pay_type;size:32;"`                        //支付方式: PIX_QRCODE=巴西PIX扫码 PKR_EASYPAISA=EasyPaisa钱包 PKR_JAZZCASH=JazzCash钱包 PHQR=菲律宾扫码 GCASH=GCASH
	AccountType     string     `json:"accountType" form:"accountType" gorm:"comment:账户类型(提现): COMPANY_BANK=对公户 PERSONAL_BANK=个人银行卡 PIX_EMAIL=PIX邮箱 PIX_PHONE=PIX手机 PIX_CPF=PIX CPF GCASH=Gcash账户;column:account_type;size:32;"` //账户类型(提现): COMPANY_BANK=对公户 PERSONAL_BANK=个人银行卡 PIX_EMAIL=PIX邮箱 PIX_PHONE=PIX手机 PIX_CPF=PIX CPF GCASH=Gcash账户
	AccountNo       string     `json:"accountNo" form:"accountNo" gorm:"comment:账号(提现用);column:account_no;size:128;"`                                                                                                           //账号(提现用)
	AccountName     string     `json:"accountName" form:"accountName" gorm:"comment:账户名(提现用);column:account_name;size:128;"`                                                                                                    //账户名(提现用)
	Content         string     `json:"content" form:"content" gorm:"comment:交易内容/备注;column:content;size:255;"`                                                                                                                  //交易内容/备注
	ClientIp        string     `json:"clientIp" form:"clientIp" gorm:"comment:客户端IP地址;column:client_ip;size:45;"`                                                                                                               //客户端IP地址
	CallbackUrl     string     `json:"callbackUrl" form:"callbackUrl" gorm:"comment:回调地址;column:callback_url;size:255;"`                                                                                                        //回调地址
	RedirectUrl     string     `json:"redirectUrl" form:"redirectUrl" gorm:"comment:跳转地址;column:redirect_url;size:255;"`                                                                                                        //跳转地址
	PayUrl          string     `json:"payUrl" form:"payUrl" gorm:"comment:支付链接;column:pay_url;size:500;"`                                                                                                                       //支付链接
	PayRaw          string     `json:"payRaw" form:"payRaw" gorm:"comment:支付原始数据;column:pay_raw;"`                                                                                                                              //支付原始数据
	ErrorMsg        string     `json:"errorMsg" form:"errorMsg" gorm:"comment:错误信息;column:error_msg;size:500;"`                                                                                                                 //错误信息
	RefCpf          string     `json:"refCpf" form:"refCpf" gorm:"comment:参考CPF;column:ref_cpf;size:32;"`                                                                                                                       //参考CPF
	RefName         string     `json:"refName" form:"refName" gorm:"comment:参考姓名;column:ref_name;size:128;"`                                                                                                                    //参考姓名
	PayEmail        string     `json:"payEmail" form:"payEmail" gorm:"comment:付款邮箱;column:pay_email;size:128;"`                                                                                                                 //付款邮箱
	PayPhone        string     `json:"payPhone" form:"payPhone" gorm:"comment:付款手机号;column:pay_phone;size:128;"`                                                                                                                //付款手机号
	PayBankCode     string     `json:"payBankCode" form:"payBankCode" gorm:"comment:付款银行编码;column:pay_bank_code;size:128;"`                                                                                                     //付款银行编码
	Type            int        `json:"type" form:"type" gorm:"comment:type;column:type;size:128;"`                                                                                                                              //付款银行编码
	CreatedAt       *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`                                                                                                                       //创建时间
	UpdatedAt       *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`                                                                                                                       //更新时间
	DeletedAt       *time.Time `json:"deletedAt" form:"deletedAt" gorm:"comment:删除时间(软删除);column:deleted_at;"`                                                                                                                  //删除时间(软删除)
}

// TableName paymentTransactions表 PaymentTransactions payment_transactions
func (PaymentTransactions) TableName() string {
	return "payment_transactions"
}
