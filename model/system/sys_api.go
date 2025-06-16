package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"comment:api"`             // api
	Description string `json:"description" gorm:"comment:api"`      // api
	ApiGroup    string `json:"apiGroup" gorm:"comment:api"`         // api
	Method      string `json:"method" gorm:"default:POST;comment:"` // :POST()|GET|PUT|DELETE
}

func (SysApi) TableName() string {
	return "sys_apis"
}

type SysIgnoreApi struct {
	global.GVA_MODEL
	Path   string `json:"path" gorm:"comment:api"`             // api
	Method string `json:"method" gorm:"default:POST;comment:"` // :POST()|GET|PUT|DELETE
	Flag   bool   `json:"flag" gorm:"-"`                       //
}

func (SysIgnoreApi) TableName() string {
	return "sys_ignore_apis"
}
