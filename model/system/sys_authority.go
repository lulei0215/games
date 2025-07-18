package system

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time       //
	UpdatedAt       time.Time       //
	DeletedAt       *time.Time      `sql:"index"`
	AuthorityId     uint            `json:"authorityId" gorm:"not null;unique;primary_key;comment:ID;size:90"` // ID
	AuthorityName   string          `json:"authorityName" gorm:"comment:"`                                     //
	ParentId        *uint           `json:"parentId" gorm:"comment:ID"`                                        // ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:;default:dashboard"` // (dashboard)
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
