// SysParams
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysParams
type SysParams struct {
	global.GVA_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;comment:;" binding:"required"`    //
	Key   string `json:"key" form:"key" gorm:"column:key;comment:;" binding:"required"`       //
	Value string `json:"value" form:"value" gorm:"column:value;comment:;" binding:"required"` //
	Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:;"`                       //
}

// TableName  SysParams sys_params
func (SysParams) TableName() string {
	return "sys_params"
}
