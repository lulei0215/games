package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserRebatesSearch struct {
	request.PageInfo
	Id                *uint      `json:"id" form:"id"`
	UserId            *uint      `json:"userId" form:"userId"`
	FromUserId        *uint      `json:"fromUserId" form:"fromUserId"`
	FromUserCode      string     `json:"fromUserCode" form:"fromUserCode"`
	RebateType        string     `json:"rebateType" form:"rebateType"`
	RebateLevel       *int       `json:"rebateLevel" form:"rebateLevel"`
	Coin              *float64   `json:"coin" form:"coin"`
	Win               *float64   `json:"win" form:"win"`
	RebateRate        *float64   `json:"rebateRate" form:"rebateRate"`
	RebateAmount      *float64   `json:"rebateAmount" form:"rebateAmount"`
	UserBalanceBefore *float64   `json:"userBalanceBefore" form:"userBalanceBefore"`
	UserBalanceAfter  *float64   `json:"userBalanceAfter" form:"userBalanceAfter"`
	SessionId         string     `json:"sessionId" form:"sessionId"`
	GameType          *int       `json:"gameType" form:"gameType"`
	Area              string     `json:"area" form:"area"`
	BetInfo           string     `json:"betInfo" form:"betInfo"`
	Status            *int       `json:"status" form:"status"`
	Remark            string     `json:"remark" form:"remark"`
	CreatedAtStart    *time.Time `json:"createdAtStart" form:"createdAtStart"`
	CreatedAtEnd      *time.Time `json:"createdAtEnd" form:"createdAtEnd"`
	UpdatedAtStart    *time.Time `json:"updatedAtStart" form:"updatedAtStart"`
	UpdatedAtEnd      *time.Time `json:"updatedAtEnd" form:"updatedAtEnd"`
}
