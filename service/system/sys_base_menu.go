package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type BaseMenuService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBaseMenu
//@description:
//@param: id float64
//@return: err error

var BaseMenuServiceApp = new(BaseMenuService)

func (baseMenuService *BaseMenuService) DeleteBaseMenu(id int) (err error) {
	err = global.GVA_DB.First(&system.SysBaseMenu{}, "parent_id = ?", id).Error
	if err == nil {
		return errors.New("")
	}
	var menu system.SysBaseMenu
	err = global.GVA_DB.First(&menu, id).Error
	if err != nil {
		return errors.New("")
	}
	err = global.GVA_DB.First(&system.SysAuthority{}, "default_router = ?", menu.Name).Error
	if err == nil {
		return errors.New("，")
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		err = tx.Delete(&system.SysBaseMenu{}, "id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Delete(&system.SysBaseMenuBtn{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&system.SysAuthorityBtn{}, "sys_menu_id = ?", id).Error
		if err != nil {
			return err
		}

		err = tx.Delete(&system.SysAuthorityMenu{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		return nil
	})

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBaseMenu
//@description:
//@param: menu model.SysBaseMenu
//@return: err error

func (baseMenuService *BaseMenuService) UpdateBaseMenu(menu system.SysBaseMenu) (err error) {
	var oldMenu system.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["transition_type"] = menu.TransitionType
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["active_name"] = menu.ActiveName
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.GVA_LOG.Debug("name")
				return errors.New("name")
			}
		}
		txErr := tx.Unscoped().Delete(&system.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		txErr = tx.Unscoped().Delete(&system.SysBaseMenuBtn{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				global.GVA_LOG.Debug(txErr.Error())
				return txErr
			}
		}

		if len(menu.MenuBtn) > 0 {
			for k := range menu.MenuBtn {
				menu.MenuBtn[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.MenuBtn).Error
			if txErr != nil {
				global.GVA_LOG.Debug(txErr.Error())
				return txErr
			}
		}

		txErr = tx.Model(&oldMenu).Updates(upDateMap).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuById
//@description: menu
//@param: id float64
//@return: menu system.SysBaseMenu, err error

func (baseMenuService *BaseMenuService) GetBaseMenuById(id int) (menu system.SysBaseMenu, err error) {
	err = global.GVA_DB.Preload("MenuBtn").Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
