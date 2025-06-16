package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Transaction 交易记录表
type Transaction struct {
	global.GVA_MODEL
	UserID        uint    `json:"userId" gorm:"comment:用户ID"`
	Amount        float64 `json:"amount" gorm:"type:decimal(20,2);comment:交易金额"`
	Type          int     `json:"type" gorm:"comment:交易类型 1:充值 2:支出 3:提现"`
	Status        int     `json:"status" gorm:"comment:状态 0:处理中 1:成功 2:失败"`
	OrderNo       string  `json:"orderNo" gorm:"size:64;comment:订单号"`
	PaymentMethod string  `json:"paymentMethod" gorm:"size:32;comment:支付方式"`
	Remark        string  `json:"remark" gorm:"size:255;comment:备注"`
	User          SysUser `json:"user" gorm:"foreignKey:UserID"`
}

// TableName 设置表名
func (Transaction) TableName() string {
	return "sys_transactions"
}
