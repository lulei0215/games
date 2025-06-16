package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Withdrawal
type Withdrawal struct {
	global.GVA_MODEL
	UserID        uint    `json:"userId" gorm:"comment:ID"`
	Amount        float64 `json:"amount" gorm:"type:decimal(20,2);comment:"`
	Status        int     `json:"status" gorm:"comment: 0: 1: 2: 3:"`
	BankName      string  `json:"bankName" gorm:"size:64;comment:"`
	BankAccount   string  `json:"bankAccount" gorm:"size:64;comment:"`
	AccountName   string  `json:"accountName" gorm:"size:64;comment:"`
	TransactionID uint    `json:"transactionId" gorm:"comment:ID"`
	Remark        string  `json:"remark" gorm:"size:255;comment:"`
	AuditRemark   string  `json:"auditRemark" gorm:"size:255;comment:"`
	User          SysUser `json:"user" gorm:"foreignKey:UserID"`
}

// TableName
func (Withdrawal) TableName() string {
	return "sys_withdrawals"
}
