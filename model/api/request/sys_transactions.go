package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SysTransactionsSearch struct {
	request.PageInfo
}
type Betting struct {
	Coin uint `json:"coin"`
	Room uint `json:"room"`
}

type Settle struct {
	SessionId int `json:"sessionid"`
	Gid       int `json:"gid"`
}

type LotteryInput struct {
	PreviousSeedHash string // 上次开奖的种子哈希（开奖前已知）
	TimeStamp        int64  // 开奖时间戳（开奖前已知）
	LuckyNumber      int64  // 开奖时间戳（开奖前已知）
	SeedString       int64  // 开奖时间戳（开奖前已知）
}

type VerifyInput struct {
	PreviousSeedHash string `json:"previous_seed_hash"` // 上次开奖的种子哈希（开奖前已知）
	TimeStamp        int64  `json:"time_stamp"`         // 开奖时间戳（开奖前已知）
	LuckyNumber      int    `json:"lucky_number"`       // 公布的幸运号码（开奖后生成）
	CurrentSeedHash  string `json:"current_seed_hash"`  // 公布的当前种子哈希（开奖后生成）
}

type DecryptRequest struct {
	Data string `json:"data" binding:"required"`
	IV   string `json:"iv" binding:"required"`
}
type RobotRequest struct {
	Limit int `json:"limit" binding:"required"`
}

type SettleRecord struct {
	SessionID string  `json:"session_id" gorm:"column:session_id"`
	UserCode  string  `json:"usercode" gorm:"column:usercode"`
	Coin      float64 `json:"coin" gorm:"column:coin"`
	BetInfo   BetInfo `json:"bet_info" gorm:"column:bet_info"` // 如果需要可定义为 map[string]interface{}
	Win       float64 `json:"win" gorm:"column:win"`
	GameType  int     `json:"gametype" gorm:"column:gametype"`
	Area      string  `json:"area" gorm:"column:area"`
	Balance   float64 `json:"balance" gorm:"column:balance"`
}
type BetInfo struct {
	Room       int `json:"room"`
	BetAmount  int `json:"betAmount"`
	TargetRoom int `json:"targetRoom"`
}
