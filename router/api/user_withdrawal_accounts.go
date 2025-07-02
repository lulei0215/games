package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserWithdrawalAccountsRouter struct{}

// InitUserWithdrawalAccountsRouter  userWithdrawalAccounts
func (s *UserWithdrawalAccountsRouter) InitUserWithdrawalAccountsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userWithdrawalAccountsRouter := Router.Group("userWithdrawalAccounts").Use(middleware.OperationRecord())
	userWithdrawalAccountsRouterWithoutRecord := Router.Group("userWithdrawalAccounts")
	userWithdrawalAccountsRouterWithoutAuth := PublicRouter.Group("userWithdrawalAccounts")
	{
		userWithdrawalAccountsRouter.POST("createUserWithdrawalAccounts", userWithdrawalAccountsApi.CreateUserWithdrawalAccounts)             // userWithdrawalAccounts
		userWithdrawalAccountsRouter.DELETE("deleteUserWithdrawalAccounts", userWithdrawalAccountsApi.DeleteUserWithdrawalAccounts)           // userWithdrawalAccounts
		userWithdrawalAccountsRouter.DELETE("deleteUserWithdrawalAccountsByIds", userWithdrawalAccountsApi.DeleteUserWithdrawalAccountsByIds) // userWithdrawalAccounts
		userWithdrawalAccountsRouter.PUT("updateUserWithdrawalAccounts", userWithdrawalAccountsApi.UpdateUserWithdrawalAccounts)              // userWithdrawalAccounts
	}
	{
		userWithdrawalAccountsRouterWithoutRecord.GET("findUserWithdrawalAccounts", userWithdrawalAccountsApi.FindUserWithdrawalAccounts)       // IDuserWithdrawalAccounts
		userWithdrawalAccountsRouterWithoutRecord.GET("getUserWithdrawalAccountsList", userWithdrawalAccountsApi.GetUserWithdrawalAccountsList) // userWithdrawalAccounts
	}
	{
		userWithdrawalAccountsRouterWithoutAuth.POST("add", userWithdrawalAccountsApi.Add)         // userWithdrawalAccounts
		userWithdrawalAccountsRouterWithoutAuth.POST("addpost", userWithdrawalAccountsApi.AddPost) // userWithdrawalAccounts
		userWithdrawalAccountsRouterWithoutAuth.POST("del", userWithdrawalAccountsApi.Del)         // userWithdrawalAccounts
		userWithdrawalAccountsRouterWithoutAuth.POST("upd", userWithdrawalAccountsApi.Upd)         // userWithdrawalAccounts
		userWithdrawalAccountsRouterWithoutAuth.POST("list", userWithdrawalAccountsApi.List)       // userWithdrawalAccounts
	}
}
