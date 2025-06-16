package system

import (
	"github.com/gin-gonic/gin"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup, RouterPublic *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	publicAutoCodeRouter := RouterPublic.Group("autoCode")
	{
		autoCodeRouter.GET("getDB", autoCodeApi.GetDB)         //
		autoCodeRouter.GET("getTables", autoCodeApi.GetTables) //
		autoCodeRouter.GET("getColumn", autoCodeApi.GetColumn) //
	}
	{
		autoCodeRouter.POST("preview", autoCodeTemplateApi.Preview)   //
		autoCodeRouter.POST("createTemp", autoCodeTemplateApi.Create) //
		autoCodeRouter.POST("addFunc", autoCodeTemplateApi.AddFunc)   //
	}
	{
		autoCodeRouter.POST("mcp", autoCodeTemplateApi.MCP)         // Mcp Tool
		autoCodeRouter.POST("mcpList", autoCodeTemplateApi.MCPList) // MCP ToolList
		autoCodeRouter.POST("mcpTest", autoCodeTemplateApi.MCPTest) // MCP
	}
	{
		autoCodeRouter.POST("getPackage", autoCodePackageApi.All)       // package
		autoCodeRouter.POST("delPackage", autoCodePackageApi.Delete)    // package
		autoCodeRouter.POST("createPackage", autoCodePackageApi.Create) // package
	}
	{
		autoCodeRouter.GET("getTemplates", autoCodePackageApi.Templates) // package
	}
	{
		autoCodeRouter.POST("pubPlug", autoCodePluginApi.Packaged)      //
		autoCodeRouter.POST("installPlugin", autoCodePluginApi.Install) //

	}
	{
		publicAutoCodeRouter.POST("llmAuto", autoCodeApi.LLMAuto)
		publicAutoCodeRouter.POST("initMenu", autoCodePluginApi.InitMenu) //
		publicAutoCodeRouter.POST("initAPI", autoCodePluginApi.InitAPI)   // API
	}
}
