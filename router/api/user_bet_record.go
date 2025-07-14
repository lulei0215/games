package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserBetRecordRouter struct {}

// InitUserBetRecordRouter  userBetRecord表 
func (s *UserBetRecordRouter) InitUserBetRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	userBetRecordRouter := Router.Group("userBetRecord").Use(middleware.OperationRecord())
	userBetRecordRouterWithoutRecord := Router.Group("userBetRecord")
	userBetRecordRouterWithoutAuth := PublicRouter.Group("userBetRecord")
	{
		userBetRecordRouter.POST("createUserBetRecord", userBetRecordApi.CreateUserBetRecord)   // userBetRecord表
		userBetRecordRouter.DELETE("deleteUserBetRecord", userBetRecordApi.DeleteUserBetRecord) // userBetRecord表
		userBetRecordRouter.DELETE("deleteUserBetRecordByIds", userBetRecordApi.DeleteUserBetRecordByIds) // userBetRecord表
		userBetRecordRouter.PUT("updateUserBetRecord", userBetRecordApi.UpdateUserBetRecord)    // userBetRecord表
	}
	{
		userBetRecordRouterWithoutRecord.GET("findUserBetRecord", userBetRecordApi.FindUserBetRecord)        // IDuserBetRecord表
		userBetRecordRouterWithoutRecord.GET("getUserBetRecordList", userBetRecordApi.GetUserBetRecordList)  // userBetRecord表
	}
	{
	    userBetRecordRouterWithoutAuth.GET("getUserBetRecordPublic", userBetRecordApi.GetUserBetRecordPublic)  // userBetRecord表
	}
}
