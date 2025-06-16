package system

type SysAuthorityBtn struct {
	AuthorityId      uint           `gorm:"comment:ID"`
	SysMenuID        uint           `gorm:"comment:ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:"`
}
