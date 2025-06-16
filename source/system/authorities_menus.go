package system

import (
	"context"

	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenuAuthority = initOrderMenu + initOrderAuthority

type initMenuAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenuAuthority, &initMenuAuthority{})
}

func (i *initMenuAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initMenuAuthority) TableCreated(ctx context.Context) bool {
	return false // always replace
}

func (i *initMenuAuthority) InitializerName() string {
	return "sys_menu_authorities"
}

func (i *initMenuAuthority) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	initAuth := &initAuthority{}
	authorities, ok := ctx.Value(initAuth.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return ctx, errors.Wrap(system.ErrMissingDependentContext, " [-] , ")
	}

	allMenus, ok := ctx.Value(new(initMenu).InitializerName()).([]sysModel.SysBaseMenu)
	if !ok {
		return next, errors.Wrap(errors.New(""), " [-] , ")
	}
	next = ctx

	// ID，
	menuMap := make(map[uint]sysModel.SysBaseMenu)
	for _, menu := range allMenus {
		menuMap[menu.ID] = menu
	}

	//
	// 1. (888) -
	if err = db.Model(&authorities[0]).Association("SysBaseMenus").Replace(allMenus); err != nil {
		return next, errors.Wrap(err, "")
	}

	// 2. (8881) -
	//
	var menu8881 []sysModel.SysBaseMenu

	// 、
	for _, menu := range allMenus {
		if menu.ParentId == 0 && (menu.Name == "dashboard" || menu.Name == "about" || menu.Name == "person" || menu.Name == "state") {
			menu8881 = append(menu8881, menu)
		}
	}

	if err = db.Model(&authorities[1]).Association("SysBaseMenus").Replace(menu8881); err != nil {
		return next, errors.Wrap(err, "")
	}

	// 3. (9528) -
	var menu9528 []sysModel.SysBaseMenu

	//
	for _, menu := range allMenus {
		if menu.ParentId == 0 {
			menu9528 = append(menu9528, menu)
		}
	}

	//  - 、
	for _, menu := range allMenus {
		parentName := ""
		if menu.ParentId > 0 && menuMap[menu.ParentId].Name != "" {
			parentName = menuMap[menu.ParentId].Name
		}

		if menu.ParentId > 0 && (parentName == "systemTools" || parentName == "example") {
			menu9528 = append(menu9528, menu)
		}
	}

	if err = db.Model(&authorities[2]).Association("SysBaseMenus").Replace(menu9528); err != nil {
		return next, errors.Wrap(err, "")
	}

	return next, nil
}

func (i *initMenuAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	auth := &sysModel.SysAuthority{}
	if ret := db.Model(auth).
		Where("authority_id = ?", 9528).Preload("SysBaseMenus").Find(auth); ret != nil {
		if ret.Error != nil {
			return false
		}
		return len(auth.SysBaseMenus) > 0
	}
	return false
}
