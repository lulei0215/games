// MonitorTransfers
package api

import (
	"time"
)

// monitorTransfers表   MonitorTransfers
type MonitorTransfers struct {
	Id           uint       `json:"id" form:"id" gorm:"primarykey;comment:ID;column:id;size:20;"`                                //ID
	To           string     `json:"to" form:"to" gorm:"comment:to;column:to;size:64;"`                                           //to
	From         string     `json:"from" form:"from" gorm:"comment:from;column:from;size:64;"`                                   //from
	Amount       float64    `json:"amount" form:"amount" gorm:"comment:amount;column:amount;size:20;"`                           //amount
	TxHash       string     `json:"txHash" form:"txHash" gorm:"comment:tx_hash;column:tx_hash;size:128;"`                        //tx_hash
	TxTime       *time.Time `json:"txTime" form:"txTime" gorm:"comment:tx_time;column:tx_time;"`                                 //tx_time
	Status       string     `json:"status" form:"status" gorm:"comment:status;column:status;size:20;"`                           //status
	Fee          float64    `json:"fee" form:"fee" gorm:"comment:fee;column:fee;size:20;"`                                       //fee
	TokenAddress string     `json:"tokenAddress" form:"tokenAddress" gorm:"comment:token_address;column:token_address;size:64;"` //token_address
	TokenSymbol  string     `json:"tokenSymbol" form:"tokenSymbol" gorm:"comment:token_symbol;column:token_symbol;size:20;"`     //token_symbol
	BlockNumber  uint       `json:"blockNumber" form:"blockNumber" gorm:"comment:block_number;column:block_number;size:19;"`     //block_number
	CreatedAt    *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:created_at;column:created_at;"`                     //created_at
	UpdatedAt    *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:updated_at;column:updated_at;"`                     //updated_at
}

// TableName monitorTransfers表 MonitorTransfers monitor_transfers
func (MonitorTransfers) TableName() string {
	return "monitor_transfers"
}
