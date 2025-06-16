// SysDictionary
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// time.Time import time
type SysDictionary struct {
	global.GVA_MODEL
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:（）"`     // （）
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:（）"`     // （）
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:"` //
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:"`       //
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

func (SysDictionary) TableName() string {
	return "sys_dictionaries"
}
