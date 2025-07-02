// UserAgentRelation
package api

import (
	"time"
)

// userAgentRelation表   UserAgentRelation
type UserAgentRelation struct {
	UserId    *int       `json:"userId" form:"userId" gorm:"primarykey;comment:id;column:user_id;size:19;"` //id
	ParentId1 *int       `json:"parentId1" form:"parentId1" gorm:"comment:parent_id_1;column:parent_id_1;"` //parent_id_1
	ParentId2 *int       `json:"parentId2" form:"parentId2" gorm:"comment:parent_id_2;column:parent_id_2;"` //parent_id_2
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                      //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                      //updatedAt字段
}

// UserAgentRelationWithUser 包含用户信息的代理关系
type UserAgentRelationWithUser struct {
	UserId    *int       `json:"userId" form:"userId" gorm:"primarykey;comment:id;column:user_id;size:19;"` //id
	ParentId1 *int       `json:"parentId1" form:"parentId1" gorm:"comment:parent_id_1;column:parent_id_1;"` //parent_id_1
	ParentId2 *int       `json:"parentId2" form:"parentId2" gorm:"comment:parent_id_2;column:parent_id_2;"` //parent_id_2
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                      //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                      //updatedAt字段
	// 用户信息
	Username string  `json:"username" gorm:"column:username;comment:用户名"`
	NickName string  `json:"nickName" gorm:"column:nick_name;comment:昵称"`
	Phone    string  `json:"phone" gorm:"column:phone;comment:手机号"`
	Email    string  `json:"email" gorm:"column:email;comment:邮箱"`
	Balance  float64 `json:"balance" gorm:"column:balance;comment:余额"`
	VipLevel uint8   `json:"vipLevel" gorm:"column:vip_level;comment:VIP等级"`
	Enable   int     `json:"enable" gorm:"column:enable;comment:状态"`
}

// TableName userAgentRelation表 UserAgentRelation user_agent_relation
func (UserAgentRelation) TableName() string {
	return "user_agent_relation"
}

// TableName userAgentRelation表 UserAgentRelationWithUser user_agent_relation
func (UserAgentRelationWithUser) TableName() string {
	return "user_agent_relation"
}
