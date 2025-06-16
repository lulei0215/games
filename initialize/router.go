package initialize

import (
	"net/http"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

//

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	sseServer := McpRun()

	// mcp
	Router.GET(global.GVA_CONFIG.MCP.SSEPath, func(c *gin.Context) {
		sseServer.SSEHandler().ServeHTTP(c.Writer, c.Request)
	})

	Router.POST(global.GVA_CONFIG.MCP.MessagePath, func(c *gin.Context) {
		sseServer.MessageHandler().ServeHTTP(c.Writer, c.Request)
	})

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// nginx， web/.env.production
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	//  npm run build。3
	// Router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist
	// Router.StaticFile("/", "./dist/index.html") //

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // https   core/server.go   Router.RunTLS("","cre/pem","key")
	// ，
	// Router.Use(middleware.Cors()) //
	// Router.Use(middleware.CorsByRules()) //
	// global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	//

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		//
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) //
		systemRouter.InitInitRouter(PublicGroup) //
	}

	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)               // api
		systemRouter.InitJwtRouter(PrivateGroup)                            // jwt
		systemRouter.InitUserRouter(PrivateGroup, PublicGroup)              //
		systemRouter.InitMenuRouter(PrivateGroup)                           // menu
		systemRouter.InitSystemRouter(PrivateGroup)                         // system
		systemRouter.InitCasbinRouter(PrivateGroup)                         //
		systemRouter.InitAutoCodeRouter(PrivateGroup, PublicGroup)          //
		systemRouter.InitAuthorityRouter(PrivateGroup)                      //
		systemRouter.InitSysDictionaryRouter(PrivateGroup)                  //
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)                //
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)             //
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)            //
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)             //
		systemRouter.InitSysExportTemplateRouter(PrivateGroup, PublicGroup) //
		systemRouter.InitSysParamsRouter(PrivateGroup, PublicGroup)         //
		exampleRouter.InitCustomerRouter(PrivateGroup)                      //
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)         //
		exampleRouter.InitAttachmentCategoryRouterRouter(PrivateGroup)      //

	}

	//
	InstallPlugin(PrivateGroup, PublicGroup, Router)

	//
	initBizRouter(PrivateGroup, PublicGroup)

	global.GVA_ROUTERS = Router.Routes()

	global.GVA_LOG.Info("router register success")
	return Router
}
