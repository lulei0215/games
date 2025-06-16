package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")

	apiPublicRouterWithoutRecord := RouterPub.Group("api")
	{
		apiRouter.GET("getApiGroups", apiRouterApi.GetApiGroups)          //
		apiRouter.GET("syncApi", apiRouterApi.SyncApi)                    // Api
		apiRouter.POST("ignoreApi", apiRouterApi.IgnoreApi)               // Api
		apiRouter.POST("enterSyncApi", apiRouterApi.EnterSyncApi)         // Api
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // Api
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // api
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // api
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // Api
	}
	{
		apiPublicRouterWithoutRecord.GET("freshCasbin", apiRouterApi.FreshCasbin) // casbin
	}
}
