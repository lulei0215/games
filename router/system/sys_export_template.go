package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysExportTemplateRouter struct {
}

// InitSysExportTemplateRouter
func (s *SysExportTemplateRouter) InitSysExportTemplateRouter(Router *gin.RouterGroup, pubRouter *gin.RouterGroup) {
	sysExportTemplateRouter := Router.Group("sysExportTemplate").Use(middleware.OperationRecord())
	sysExportTemplateRouterWithoutRecord := Router.Group("sysExportTemplate")
	sysExportTemplateRouterWithoutAuth := pubRouter.Group("sysExportTemplate")

	{
		sysExportTemplateRouter.POST("createSysExportTemplate", exportTemplateApi.CreateSysExportTemplate)             //
		sysExportTemplateRouter.DELETE("deleteSysExportTemplate", exportTemplateApi.DeleteSysExportTemplate)           //
		sysExportTemplateRouter.DELETE("deleteSysExportTemplateByIds", exportTemplateApi.DeleteSysExportTemplateByIds) //
		sysExportTemplateRouter.PUT("updateSysExportTemplate", exportTemplateApi.UpdateSysExportTemplate)              //
		sysExportTemplateRouter.POST("importExcel", exportTemplateApi.ImportExcel)                                     // excel
	}
	{
		sysExportTemplateRouterWithoutRecord.GET("findSysExportTemplate", exportTemplateApi.FindSysExportTemplate)       // ID
		sysExportTemplateRouterWithoutRecord.GET("getSysExportTemplateList", exportTemplateApi.GetSysExportTemplateList) //
		sysExportTemplateRouterWithoutRecord.GET("exportExcel", exportTemplateApi.ExportExcel)                           // token
		sysExportTemplateRouterWithoutRecord.GET("exportTemplate", exportTemplateApi.ExportTemplate)                     //
	}
	{
		sysExportTemplateRouterWithoutAuth.GET("exportExcelByToken", exportTemplateApi.ExportExcelByToken)       // token
		sysExportTemplateRouterWithoutAuth.GET("exportTemplateByToken", exportTemplateApi.ExportTemplateByToken) // token
	}
}
