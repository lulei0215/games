package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysBaseMenu struct {
	global.GVA_MODEL
	MenuLevel     uint                                   `json:"-"`
	ParentId      uint                                   `json:"parentId" gorm:"comment:ID"` // ID
	Path          string                                 `json:"path" gorm:"comment:path"`   // path
	Name          string                                 `json:"name" gorm:"comment:name"`   // name
	Hidden        bool                                   `json:"hidden" gorm:"comment:"`     //
	Component     string                                 `json:"component" gorm:"comment:"`  //
	Sort          int                                    `json:"sort" gorm:"comment:"`       //
	Meta          `json:"meta" gorm:"embedded;comment:"` //
	SysAuthoritys []SysAuthority                         `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu                          `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter                 `json:"parameters"`
	MenuBtn       []SysBaseMenuBtn                       `json:"menuBtn"`
}

type Meta struct {
	ActiveName     string `json:"activeName" gorm:"comment:"`
	KeepAlive      bool   `json:"keepAlive" gorm:"comment:"`      //
	DefaultMenu    bool   `json:"defaultMenu" gorm:"comment:（）"`  // （）
	Title          string `json:"title" gorm:"comment:"`          //
	Icon           string `json:"icon" gorm:"comment:"`           //
	CloseTab       bool   `json:"closeTab" gorm:"comment:tab"`    // tab
	TransitionType string `json:"transitionType" gorm:"comment:"` //
}

type SysBaseMenuParameter struct {
	global.GVA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:paramsquery"` // paramsquery
	Key           string `json:"key" gorm:"comment:key"`          // key
	Value         string `json:"value" gorm:"comment:"`           //
}

func (SysBaseMenu) TableName() string {
	return "sys_base_menus"
}
