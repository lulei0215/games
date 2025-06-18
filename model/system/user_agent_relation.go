// UserAgentRelation
package system

import (
	"time"
)

// userAgentRelation表   UserAgentRelation
type UserAgentRelation struct {
	UserId    int        `json:"userId" form:"userId" gorm:"primarykey;comment:id;column:user_id;size:19;"`         //id
	ParentId1 int        `json:"parentId1" form:"parentId1" gorm:"comment:parent_id_1;column:parent_id_1;size:19;"` //parent_id_1
	ParentId2 int        `json:"parentId2" form:"parentId2" gorm:"comment:parent_id_2;column:parent_id_2;size:19;"` //parent_id_2
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                              //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                              //updatedAt字段
}

// TableName userAgentRelation表 UserAgentRelation user_agent_relation
func (UserAgentRelation) TableName() string {
	return "user_agent_relation"
}
