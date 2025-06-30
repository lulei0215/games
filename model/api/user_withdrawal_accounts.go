// UserWithdrawalAccounts
package api

import (
	"time"
)

// userWithdrawalAccounts表   UserWithdrawalAccounts
type UserWithdrawalAccounts struct {
	Id            uint       `json:"id" form:"id" gorm:"primarykey;comment:主键ID;column:id;size:20;"`                                                                           //ID
	UserId        uint       `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;size:20;"`                                                                         //userID
	AccountName   string     `json:"accountName" form:"accountName" gorm:"comment:真实姓名;column:account_name;size:100;"`                                                         //name
	AccountType   string     `json:"accountType" form:"accountType" gorm:"comment:账户类型：PIX_PHONE=手机号，PIX_EMAIL=邮箱，PIX_CPF=CPF号码，PIX_CNPJ=CNPJ号码;column:account_type;size:20;"` //PIX_PHONE，PIX_EMAIL，PIX_CPF=CPF，PIX_CNPJ=CNPJ
	AccountNumber string     `json:"accountNumber" form:"accountNumber" gorm:"comment:账户号码（手机号/邮箱/PIX账户等）;column:account_number;size:100;"`                                    //账户号码
	CpfNumber     string     `json:"cpfNumber" form:"cpfNumber" gorm:"comment:11位CPF号码;column:cpf_number;size:14;"`                                                            //CP
	IsDefault     int        `json:"isDefault" form:"isDefault" gorm:"comment:是否为默认账户：0=否，1=是;column:is_default;"`                                                             //default
	Status        int        `json:"status" form:"status" gorm:"comment:账户状态：0=禁用，1=启用;column:status;"`                                                                        //0 not 1ok
	CreatedAt     *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`                                                                        //createtime
	UpdatedAt     *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`                                                                        //updatetime
	DeletedAt     *time.Time `json:"deletedAt" form:"deletedAt" gorm:"comment:删除时间（软删除）;column:deleted_at;"`                                                                   //deletetime
}

// TableName userWithdrawalAccounts表 UserWithdrawalAccounts user_withdrawal_accounts
func (UserWithdrawalAccounts) TableName() string {
	return "user_withdrawal_accounts"
}
