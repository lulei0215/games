package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysParamsRouter struct{}

// InitSysParamsRouter
func (s *SysParamsRouter) InitSysParamsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysParamsRouter := Router.Group("sysParams").Use(middleware.OperationRecord())
	sysParamsRouterWithoutRecord := Router.Group("sysParams")
	{
		sysParamsRouter.POST("createSysParams", sysParamsApi.CreateSysParams)             //
		sysParamsRouter.DELETE("deleteSysParams", sysParamsApi.DeleteSysParams)           //
		sysParamsRouter.DELETE("deleteSysParamsByIds", sysParamsApi.DeleteSysParamsByIds) //
		sysParamsRouter.PUT("updateSysParams", sysParamsApi.UpdateSysParams)              //
	}
	{
		sysParamsRouterWithoutRecord.GET("findSysParams", sysParamsApi.FindSysParams)       // ID
		sysParamsRouterWithoutRecord.GET("getSysParamsList", sysParamsApi.GetSysParamsList) //
		sysParamsRouterWithoutRecord.GET("getSysParam", sysParamsApi.GetSysParam)           // Key
	}
}
