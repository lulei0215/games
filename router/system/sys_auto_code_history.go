package system

import (
	"github.com/gin-gonic/gin"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	{
		autoCodeHistoryRouter.POST("getMeta", autocodeHistoryApi.First)         // idmeta
		autoCodeHistoryRouter.POST("rollback", autocodeHistoryApi.RollBack)     //
		autoCodeHistoryRouter.POST("delSysHistory", autocodeHistoryApi.Delete)  //
		autoCodeHistoryRouter.POST("getSysHistory", autocodeHistoryApi.GetList) //
	}
}
