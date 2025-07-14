package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRebatesRouter struct{}

// InitUserRebatesRouter  userRebates表
func (s *UserRebatesRouter) InitUserRebatesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userRebatesRouter := Router.Group("userRebates").Use(middleware.OperationRecord())
	userRebatesRouterWithoutRecord := Router.Group("userRebates")
	userRebatesRouterWithoutAuth := PublicRouter.Group("userRebates")
	{
		userRebatesRouter.POST("createUserRebates", userRebatesApi.CreateUserRebates)             // userRebates表
		userRebatesRouter.DELETE("deleteUserRebates", userRebatesApi.DeleteUserRebates)           // userRebates表
		userRebatesRouter.DELETE("deleteUserRebatesByIds", userRebatesApi.DeleteUserRebatesByIds) // userRebates表
		userRebatesRouter.PUT("updateUserRebates", userRebatesApi.UpdateUserRebates)              // userRebates表
	}
	{
		userRebatesRouterWithoutRecord.GET("findUserRebates", userRebatesApi.FindUserRebates)       // IDuserRebates表
		userRebatesRouterWithoutRecord.GET("getUserRebatesList", userRebatesApi.GetUserRebatesList) // userRebates表
	}
	{
		userRebatesRouterWithoutAuth.GET("getUserRebatesPublic", userRebatesApi.GetUserRebatesPublic) // userRebates表
		userRebatesRouterWithoutAuth.GET("getMyRebatesList", userRebatesApi.GetMyRebatesList)         // userRebates表

	}
}
