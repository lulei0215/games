// UserWithdrawalAccounts
package api

import (
	"time"
)

// UserWithdrawalAccounts 用户提现账户信息表
type UserWithdrawalAccounts struct {
	Id            uint       `json:"id" form:"id" gorm:"primarykey;comment:主键ID;column:id;size:20;"`
	UserId        uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:20;index:idx_user_id;"`
	FirstName     string     `json:"firstName" form:"firstName" gorm:"comment:名;column:first_name;size:255;"`
	LastName      string     `json:"lastName" form:"lastName" gorm:"comment:姓;column:last_name;size:255;"`
	AccountName   string     `json:"accountName" form:"accountName" gorm:"comment:真实姓名;column:account_name;size:100;not null;"`
	AccountType   string     `json:"accountType" form:"accountType" gorm:"comment:账户类型：PIX_PHONE=手机号，PIX_EMAIL=邮箱，PIX_CPF=CPF号码，PIX_CNPJ=CNPJ号码;column:account_type;size:20;"`
	AccountNumber string     `json:"accountNumber" form:"accountNumber" gorm:"comment:账户号码（手机号/邮箱/PIX账户等）;column:account_number;size:100;index:idx_account_number;"`
	Email         string     `json:"email" form:"email" gorm:"comment:邮箱地址;column:email;size:255;"`
	Phone         string     `json:"phone" form:"phone" gorm:"comment:手机号码;column:phone;size:255;"`
	BankCode      string     `json:"bankCode" form:"bankCode" gorm:"comment:bankCode;column:bankCode;size:255;"`
	CnpjNumber    string     `json:"cnpjNumber" form:"cnpjNumber" gorm:"comment:CNPJ号码;column:cnpj_number;size:255;"`
	CpfNumber     string     `json:"cpfNumber" form:"cpfNumber" gorm:"comment:11位CPF号码;column:cpf_number;size:14;index:idx_cpf_number;"`
	IsDefault     int        `json:"isDefault" form:"isDefault" gorm:"comment:是否为默认账户：0=否，1=是;column:is_default;index:idx_user_default;"`
	Status        int        `json:"status" form:"status" gorm:"comment:账户状态：0=禁用，1=启用;column:status;"`
	CreatedAt     *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`
	UpdatedAt     *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`
	DeletedAt     *time.Time `json:"deletedAt" form:"deletedAt" gorm:"comment:删除时间（软删除）;column:deleted_at;index;"`
}

// TableName userWithdrawalAccounts表 UserWithdrawalAccounts user_withdrawal_accounts
func (UserWithdrawalAccounts) TableName() string {
	return "user_withdrawal_accounts"
}

// 常量定义
const (
	// 账户类型
	AccountTypePIXPhone = "PIX_PHONE" // 手机号
	AccountTypePIXEmail = "PIX_EMAIL" // 邮箱
	AccountTypePIXCPF   = "PIX_CPF"   // CPF号码
	AccountTypePIXCNPJ  = "PIX_CNPJ"  // CNPJ号码
	AccountTypePIXEVP   = "PIX_EVP"   // EVP账号
	AccountTypeBRBANK   = "BRBANK"    // 银行账户

	// 状态
	StatusDisabled = 0 // 禁用
	StatusEnabled  = 1 // 启用

	// 默认账户
	IsDefaultNo  = 0 // 非默认
	IsDefaultYes = 1 // 默认账户
)

// 请求结构体
type AddWithdrawAccountRequest struct {
	// 银行编码 - 巴西：PIX、PIXN、TED、PICPAY
	BankCode string `json:"bankCode" form:"bankCode" validate:"required"`

	// 收款人姓名（全名）
	BankAcctName string `json:"bankAcctName" form:"bankAcctName" validate:"required"`

	// 收款人FirstName
	BankFirstName string `json:"bankFirstName" form:"bankFirstName" validate:"required"`

	// 收款人LastName
	BankLastName string `json:"bankLastName" form:"bankLastName" validate:"required"`

	// 收款人账号
	BankAcctNo string `json:"bankAcctNo" form:"bankAcctNo" validate:"required"`

	// 收款人手机号
	AccPhone string `json:"accPhone" form:"accPhone" validate:"required"`

	// 证件号码/税号
	IdentityNo string `json:"identityNo" form:"identityNo" validate:"required"`

	// 证件类型/收款类型编码 - 巴西：CPF、CNPJ、PHONE、EMAIL、EVP、BRBANK
	IdentityType string `json:"identityType" form:"identityType" validate:"required"`
}

// ToUserWithdrawalAccounts 将请求转换为数据库模型
func (req *AddWithdrawAccountRequest) ToUserWithdrawalAccounts(userId uint) *UserWithdrawalAccounts {
	now := time.Now()

	// 根据证件类型设置账户类型
	accountType := req.getAccountType()

	// 根据证件类型设置对应字段
	account := &UserWithdrawalAccounts{
		UserId:        userId,
		FirstName:     req.BankFirstName,
		LastName:      req.BankLastName,
		AccountName:   req.BankAcctName,
		AccountType:   accountType,
		AccountNumber: req.BankAcctNo,
		Status:        StatusEnabled, // 默认启用
		IsDefault:     IsDefaultNo,   // 默认非默认账户
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	// 根据证件类型设置对应字段
	switch req.IdentityType {
	case "CPF":
		account.CpfNumber = req.IdentityNo
	case "CNPJ":
		account.CnpjNumber = req.IdentityNo
	case "PHONE":
		account.Phone = req.AccPhone
	case "EMAIL":
		account.Email = req.BankAcctNo
	}

	return account
}

// getAccountType 根据证件类型获取账户类型
func (req *AddWithdrawAccountRequest) getAccountType() string {
	switch req.IdentityType {
	case "CPF":
		return AccountTypePIXCPF
	case "CNPJ":
		return AccountTypePIXCNPJ
	case "PHONE":
		return AccountTypePIXPhone
	case "EMAIL":
		return AccountTypePIXEmail
	case "EVP":
		return AccountTypePIXEVP
	case "BRBANK":
		return AccountTypeBRBANK
	default:
		return req.IdentityType
	}
}

// 响应结构体
type AddWithdrawAccountResponse struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg"`
	Data *UserWithdrawalAccounts `json:"data"`
}
