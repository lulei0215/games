package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)           //
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) //	menu
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     //
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     //
	}
	{
		menuRouterWithoutRecord.POST("getMenu", authorityMenuApi.GetMenu)                   //
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)           // menu
		menuRouterWithoutRecord.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   //
		menuRouterWithoutRecord.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // menu
		menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // id
	}
	return menuRouter
}
