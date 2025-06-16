package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", dictionaryDetailApi.CreateSysDictionaryDetail)   // SysDictionaryDetail
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", dictionaryDetailApi.DeleteSysDictionaryDetail) // SysDictionaryDetail
		dictionaryDetailRouter.PUT("updateSysDictionaryDetail", dictionaryDetailApi.UpdateSysDictionaryDetail)    // SysDictionaryDetail
	}
	{
		dictionaryDetailRouterWithoutRecord.GET("findSysDictionaryDetail", dictionaryDetailApi.FindSysDictionaryDetail)       // IDSysDictionaryDetail
		dictionaryDetailRouterWithoutRecord.GET("getSysDictionaryDetailList", dictionaryDetailApi.GetSysDictionaryDetailList) // SysDictionaryDetail
	}
}
