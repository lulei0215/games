package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	{
		sysDictionaryRouter.POST("createSysDictionary", dictionaryApi.CreateSysDictionary)   // SysDictionary
		sysDictionaryRouter.DELETE("deleteSysDictionary", dictionaryApi.DeleteSysDictionary) // SysDictionary
		sysDictionaryRouter.PUT("updateSysDictionary", dictionaryApi.UpdateSysDictionary)    // SysDictionary
	}
	{
		sysDictionaryRouterWithoutRecord.GET("findSysDictionary", dictionaryApi.FindSysDictionary)       // IDSysDictionary
		sysDictionaryRouterWithoutRecord.GET("getSysDictionaryList", dictionaryApi.GetSysDictionaryList) // SysDictionary
	}
}
