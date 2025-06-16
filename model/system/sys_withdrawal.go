package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Withdrawal 提现表
type Withdrawal struct {
	global.GVA_MODEL
	UserID        uint    `json:"userId" gorm:"comment:用户ID"`
	Amount        float64 `json:"amount" gorm:"type:decimal(20,2);comment:提现金额"`
	Status        int     `json:"status" gorm:"comment:状态 0:待审核 1:已通过 2:已拒绝 3:已打款"`
	BankName      string  `json:"bankName" gorm:"size:64;comment:银行名称"`
	BankAccount   string  `json:"bankAccount" gorm:"size:64;comment:银行账号"`
	AccountName   string  `json:"accountName" gorm:"size:64;comment:开户名"`
	TransactionID uint    `json:"transactionId" gorm:"comment:关联的交易记录ID"`
	Remark        string  `json:"remark" gorm:"size:255;comment:备注"`
	AuditRemark   string  `json:"auditRemark" gorm:"size:255;comment:审核备注"`
	User          SysUser `json:"user" gorm:"foreignKey:UserID"`
}

// TableName 设置表名
func (Withdrawal) TableName() string {
	return "sys_withdrawals"
}
