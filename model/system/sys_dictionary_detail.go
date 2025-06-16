// SysDictionaryDetail
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// time.Time import time
type SysDictionaryDetail struct {
	global.GVA_MODEL
	Label           string `json:"label" form:"label" gorm:"column:label;comment:"`                                 //
	Value           string `json:"value" form:"value" gorm:"column:value;comment:"`                                 //
	Extend          string `json:"extend" form:"extend" gorm:"column:extend;comment:"`                              //
	Status          *bool  `json:"status" form:"status" gorm:"column:status;comment:"`                              //
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:"`                                    //
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id;comment:"` //
}

func (SysDictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}
