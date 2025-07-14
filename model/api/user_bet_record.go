// UserBetRecord
package api

import (
	"time"

	"gorm.io/datatypes"
)

// userBetRecord表   UserBetRecord
type UserBetRecord struct {
	Id        int            `json:"id" form:"id" gorm:"primarykey;comment:Primary Key ID;column:id;size:20;"`                                    //Primary Key ID
	SessionId string         `json:"sessionId" form:"sessionId" gorm:"comment:Session/Issue Number;column:session_id;size:64;"`                   //Session/Issue Number
	Usercode  string         `json:"usercode" form:"usercode" gorm:"comment:User Unique Identifier;column:usercode;size:64;"`                     //User Unique Identifier
	Coin      float64        `json:"coin" form:"coin" gorm:"comment:Bet Amount;column:coin;size:18;"`                                             //Bet Amount
	BetInfo   datatypes.JSON `json:"betInfo" form:"betInfo" gorm:"comment:Bet Details (Struct Array);column:bet_info;" swaggertype:"object"`      //Bet Details (Struct Array)
	Win       float64        `json:"win" form:"win" gorm:"comment:Winning Amount;column:win;size:18;"`                                            //Winning Amount
	Gametype  int            `json:"gametype" form:"gametype" gorm:"comment:Game Type;column:gametype;size:10;"`                                  //Game Type
	Area      string         `json:"area" form:"area" gorm:"comment:Bet Area;column:area;size:32;"`                                               //Bet Area
	Balance   float64        `json:"balance" form:"balance" gorm:"comment:Balance After Settlement;column:balance;size:18;"`                      //Balance After Settlement
	Cards     datatypes.JSON `json:"cards" form:"cards" gorm:"comment:Cards (2D Array, Struct);column:cards;" swaggertype:"object"`               //Cards (2D Array, Struct)
	Result    datatypes.JSON `json:"result" form:"result" gorm:"comment:Result (Struct Array);column:result;" swaggertype:"object"`               //Result (Struct Array)
	RoleCards datatypes.JSON `json:"roleCards" form:"roleCards" gorm:"comment:Role Cards (Struct Array);column:role_cards;" swaggertype:"object"` //Role Cards (Struct Array)
	CreatedAt *time.Time     `json:"createdAt" form:"createdAt" gorm:"comment:Creation Time;column:created_at;"`                                  //Creation Time
	UpdatedAt *time.Time     `json:"updatedAt" form:"updatedAt" gorm:"comment:Update Time;column:updated_at;"`                                    //Update Time
}

// TableName userBetRecord表 UserBetRecord user_bet_record
func (UserBetRecord) TableName() string {
	return "user_bet_record"
}
