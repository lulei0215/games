package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Transaction
type Transaction struct {
	global.GVA_MODEL
	UserID        uint    `json:"userId" gorm:"comment:ID"`
	Amount        float64 `json:"amount" gorm:"type:decimal(20,2);comment:"`
	Type          int     `json:"type" gorm:"comment: 1: 2: 3:"`
	Status        int     `json:"status" gorm:"comment: 0: 1: 2:"`
	OrderNo       string  `json:"orderNo" gorm:"size:64;comment:"`
	PaymentMethod string  `json:"paymentMethod" gorm:"size:32;comment:"`
	Remark        string  `json:"remark" gorm:"size:255;comment:"`
	User          SysUser `json:"user" gorm:"foreignKey:UserID"`
}

// TableName
func (Transaction) TableName() string {
	return "sys_transactions"
}
