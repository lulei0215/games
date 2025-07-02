package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserAgentRelationRouter struct{}

// InitUserAgentRelationRouter  userAgentRelation表
func (s *UserAgentRelationRouter) InitUserAgentRelationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userAgentRelationRouter := Router.Group("userAgentRelation").Use(middleware.OperationRecord())
	userAgentRelationRouterWithoutRecord := Router.Group("userAgentRelation")
	userAgentRelationRouterWithoutAuth := PublicRouter.Group("userAgentRelation")
	{
		userAgentRelationRouter.POST("createUserAgentRelation", userAgentRelationApi.CreateUserAgentRelation)             // userAgentRelation表
		userAgentRelationRouter.DELETE("deleteUserAgentRelation", userAgentRelationApi.DeleteUserAgentRelation)           // userAgentRelation表
		userAgentRelationRouter.DELETE("deleteUserAgentRelationByIds", userAgentRelationApi.DeleteUserAgentRelationByIds) // userAgentRelation表
		userAgentRelationRouter.PUT("updateUserAgentRelation", userAgentRelationApi.UpdateUserAgentRelation)              // userAgentRelation表
	}
	{
		userAgentRelationRouterWithoutRecord.GET("findUserAgentRelation", userAgentRelationApi.FindUserAgentRelation)       // IDuserAgentRelation表
		userAgentRelationRouterWithoutRecord.GET("getUserAgentRelationList", userAgentRelationApi.GetUserAgentRelationList) // userAgentRelation表
	}
	{
		userAgentRelationRouterWithoutAuth.GET("getUserAgentRelationPublic", userAgentRelationApi.GetUserAgentRelationPublic) // userAgentRelation表
		userAgentRelationRouterWithoutAuth.GET("list", userAgentRelationApi.GetList)                                          // userAgentRelation表
	}
}
