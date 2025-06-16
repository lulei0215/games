package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName       string         `json:"customerName" form:"customerName" gorm:"comment:"`               //
	CustomerPhoneData  string         `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:"`     //
	SysUserID          uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:ID"`                   // ID
	SysUserAuthorityID uint           `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:ID"` // ID
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:"`                         //
}
