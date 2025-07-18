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

type BetInfoData struct {
	Room       int         `json:"room"`
	BetAmount  int         `json:"betAmount"`
	TargetRoom interface{} `json:"targetRoom"` // 支持单个数字或数组
}

// 新增卡片相关结构体
type Card struct {
	Suit  int `json:"suit"`
	Value int `json:"value"`
}

type RankInfo struct {
	Rank   int    `json:"rank"`
	Type   string `json:"type"`
	Values []int  `json:"values"`
}

type RoleCard struct {
	Cards    []Card   `json:"cards"`
	RankInfo RankInfo `json:"rank_info"`
}

type RoleCardResult struct {
	Card   RoleCard `json:"card"`
	Result string   `json:"result"`
}

type SettleRecord struct {
	SessionID string        `json:"session_id" gorm:"column:session_id"`
	UserCode  string        `json:"usercode" gorm:"column:usercode"`
	Coin      float64       `json:"coin" gorm:"column:coin"`
	BetInfo   []BetInfoData `json:"bet_info" gorm:"column:bet_info"`
	Win       float64       `json:"win" gorm:"column:win"`
	GameType  int           `json:"gametype" gorm:"column:gametype"`
	Area      string        `json:"area" gorm:"column:area"`
	Balance   float64       `json:"balance" gorm:"column:balance"`
	// 新增字段
	Cards     [][]Card         `json:"cards,omitempty" gorm:"column:cards"`           // 可选字段，用于新数据结构
	Result    []Card           `json:"result,omitempty" gorm:"column:result"`         // 可选字段，用于新数据结构
	RoleCards []RoleCardResult `json:"role_cards,omitempty" gorm:"column:role_cards"` // 可选字段，用于新数据结构
}

type BetInfo struct {
	Room       int `json:"room"`
	BetAmount  int `json:"betAmount"`
	TargetRoom int `json:"targetRoom"`
}

type MonitorTransfer struct {
	To     string  `json:"to"  gorm:"column:to"`
	From   string  `json:"from"  gorm:"column:from"`
	Amount float64 `json:"amount"  gorm:"column:amount"`
	TxHash string  `json:"tx_hash"  gorm:"column:tx_hash"`
	TxTime string  `json:"tx_time"  gorm:"column:tx_time"`
	Status string  `json:"status"  gorm:"column:status"`
	Fee    float64 `json:"fee"  gorm:"column:fee"`
}

type MonitorTransferApi struct {
	To       string  `json:"to"  gorm:"column:to"`
	Password string  `json:"password"  gorm:"column:password"`
	Amount   float64 `json:"amount"  gorm:"column:amount"`
	TotpCode string  `json:"totpcode"  gorm:"column:totpcode"`
}

type SettleList struct {
	List      []SettleRecord   `json:"list"`
	Timestamp string           `json:"timestamp" gorm:"column:timestamp"`
	Sign      string           `json:"sign" gorm:"column:sign"`
	Cards     [][]Card         `json:"cards,omitempty" gorm:"column:cards"`           // 可选字段，用于新数据结构
	Result    []Card           `json:"result,omitempty" gorm:"column:result"`         // 可选字段，用于新数据结构
	RoleCards []RoleCardResult `json:"role_cards,omitempty" gorm:"column:role_cards"` // 可选字段，用于新数据结构
}
