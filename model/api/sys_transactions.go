// SysTransactions
package api

import (
	"time"
)

// sysTransactions   SysTransactions
type SysTransactions struct {
	Id            *int       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:20;"`                                    //ID
	CreatedAt     *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                                            //createdAt
	UpdatedAt     *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                                            //updatedAt
	DeletedAt     *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`                                            //deletedAt
	UserId        *int       `json:"userId" form:"userId" gorm:"comment:ID;column:user_id;size:20;"`                                  //ID
	Amount        *float64   `json:"amount" form:"amount" gorm:"comment:amount;column:amount;size:20;"`                               //amount
	Type          *int       `json:"type" form:"type" gorm:"comment:type 1:chong 2:chu 3:withdrawal;column:type;size:10;"`            //type 1:chong 2:chu 3:withdrawal
	Status        *int       `json:"status" form:"status" gorm:"comment:status 0:ing 1:ok 2:fail;column:status;size:10;"`             //status 0:ing 1:ok 2:fail
	OrderNo       *string    `json:"orderNo" form:"orderNo" gorm:"comment:order_no;column:order_no;size:64;"`                         //order_no
	PaymentMethod *string    `json:"paymentMethod" form:"paymentMethod" gorm:"comment:payment_method;column:payment_method;size:32;"` //payment_method
	Remark        *string    `json:"remark" form:"remark" gorm:"comment:remark;column:remark;size:255;"`                              //remark
}

// TableName sysTransactions SysTransactions sys_transactions
func (SysTransactions) TableName() string {
	return "sys_transactions"
}
