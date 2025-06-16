package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Info = new(info)

type info struct{}

// Init
func (r *info) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("info").Use(middleware.OperationRecord())
		group.POST("createInfo", apiInfo.CreateInfo)             //
		group.DELETE("deleteInfo", apiInfo.DeleteInfo)           //
		group.DELETE("deleteInfoByIds", apiInfo.DeleteInfoByIds) //
		group.PUT("updateInfo", apiInfo.UpdateInfo)              //
	}
	{
		group := private.Group("info")
		group.GET("findInfo", apiInfo.FindInfo)       // ID
		group.GET("getInfoList", apiInfo.GetInfoList) //
	}
	{
		group := public.Group("info")
		group.GET("getInfoDataSource", apiInfo.GetInfoDataSource) //
		group.GET("getInfoPublic", apiInfo.GetInfoPublic)         //
	}
}
