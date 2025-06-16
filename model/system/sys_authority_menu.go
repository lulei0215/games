package system

type SysMenu struct {
	SysBaseMenu
	MenuId      uint                   `json:"menuId" gorm:"comment:ID"`
	AuthorityId uint                   `json:"-" gorm:"comment:ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}

type SysAuthorityMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:ID;column:sys_base_menu_id"`
	AuthorityId string `json:"-" gorm:"comment:ID;column:sys_authority_authority_id"`
}

func (s SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
