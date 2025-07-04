// UserRebates
package api

import (
	"time"

	"gorm.io/datatypes"
)

// userRebates表   UserRebates
type UserRebates struct {
	Id                uint           `json:"id" form:"id" gorm:"primarykey;comment:主键ID;column:id;size:20;"`                                                    //主键ID
	UserId            uint           `json:"userId" form:"userId" gorm:"comment:user_id;column:user_id;size:20;"`                                               //user_id
	FromUserId        uint           `json:"fromUserId" form:"fromUserId" gorm:"comment:from_user_id;column:from_user_id;size:20;"`                             //from_user_id
	FromUserCode      string         `json:"fromUserCode" form:"fromUserCode" gorm:"comment:from_user_code;column:from_user_code;size:50;"`                     //from_user_code
	RebateType        string         `json:"rebateType" form:"rebateType" gorm:"comment:rebate_type;column:rebate_type;size:20;"`                               //rebate_type
	RebateLevel       int            `json:"rebateLevel" form:"rebateLevel" gorm:"comment:rebate_level;column:rebate_level;"`                                   //rebate_level
	Coin              float64        `json:"coin" form:"coin" gorm:"comment:coin;column:coin;size:10;"`                                                         //coin
	Win               float64        `json:"win" form:"win" gorm:"comment:win;column:win;size:10;"`                                                             //win
	RebateRate        float64        `json:"rebateRate" form:"rebateRate" gorm:"comment:rebate_rate;column:rebate_rate;size:5;"`                                //rebate_rate
	RebateAmount      float64        `json:"rebateAmount" form:"rebateAmount" gorm:"comment:rebate_amount;column:rebate_amount;size:10;"`                       //rebate_amount
	UserBalanceBefore float64        `json:"userBalanceBefore" form:"userBalanceBefore" gorm:"comment:user_balance_before;column:user_balance_before;size:10;"` //user_balance_before
	UserBalanceAfter  float64        `json:"userBalanceAfter" form:"userBalanceAfter" gorm:"comment:user_balance_after;column:user_balance_after;size:10;"`     //user_balance_after
	SessionId         string         `json:"sessionId" form:"sessionId" gorm:"comment:session_id;column:session_id;size:100;"`                                  //session_id
	GameType          int            `json:"gameType" form:"gameType" gorm:"comment:game_type;column:game_type;size:10;"`                                       //game_type
	Area              string         `json:"area" form:"area" gorm:"comment:area;column:area;size:50;"`                                                         //area
	BetInfo           datatypes.JSON `json:"betInfo" form:"betInfo" gorm:"comment:bet_info_json;column:bet_info;" swaggertype:"object"`                         //bet_info_json
	Status            int            `json:"status" form:"status" gorm:"comment:status;column:status;"`                                                         //status
	Remark            string         `json:"remark" form:"remark" gorm:"comment:remark;column:remark;size:255;"`                                                //remark
	CreatedAt         *time.Time     `json:"createdAt" form:"createdAt" gorm:"comment:created_at;column:created_at;"`                                           //created_at
	UpdatedAt         *time.Time     `json:"updatedAt" form:"updatedAt" gorm:"comment:updated_at;column:updated_at;"`                                           //updated_at
}

// TableName userRebates表 UserRebates user_rebates
func (UserRebates) TableName() string {
	return "user_rebates"
}
