package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysAutoCodePackage struct {
	global.GVA_MODEL
	Desc        string `json:"desc" gorm:"comment:"`
	Label       string `json:"label" gorm:"comment:"`
	Template    string `json:"template"  gorm:"comment:"`
	PackageName string `json:"packageName" gorm:"comment:"`
	Module      string `json:"-" example:""`
}

func (s *SysAutoCodePackage) TableName() string {
	return "sys_auto_code_packages"
}
