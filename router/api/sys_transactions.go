package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysTransactionsRouter struct{}

// InitSysTransactionsRouter  sysTransactions
func (s *SysTransactionsRouter) InitSysTransactionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysTransactionsRouter := Router.Group("sysTransactions").Use(middleware.OperationRecord())
	sysTransactionsRouterWithoutRecord := Router.Group("sysTransactions")
	sysTransactionsRouterWithoutAuth := PublicRouter.Group("sysTransactions")
	{
		sysTransactionsRouter.POST("createSysTransactions", sysTransactionsApi.CreateSysTransactions)             // sysTransactions
		sysTransactionsRouter.DELETE("deleteSysTransactions", sysTransactionsApi.DeleteSysTransactions)           // sysTransactions
		sysTransactionsRouter.DELETE("deleteSysTransactionsByIds", sysTransactionsApi.DeleteSysTransactionsByIds) // sysTransactions
		sysTransactionsRouter.PUT("updateSysTransactions", sysTransactionsApi.UpdateSysTransactions)              // sysTransactions
	}
	{
		sysTransactionsRouterWithoutRecord.GET("findSysTransactions", sysTransactionsApi.FindSysTransactions)       // IDsysTransactions
		sysTransactionsRouterWithoutRecord.GET("getSysTransactionsList", sysTransactionsApi.GetSysTransactionsList) // sysTransactions
	}
	{
		sysTransactionsRouterWithoutAuth.GET("getSysTransactionsPublic", sysTransactionsApi.GetSysTransactionsPublic) // sysTransactions
		sysTransactionsRouterWithoutAuth.GET("get", sysTransactionsApi.Get)                                           // sysTransactions
		sysTransactionsRouterWithoutAuth.POST("settle", sysTransactionsApi.Settle)                                    // sysTransactions
		sysTransactionsRouterWithoutAuth.POST("lottery", sysTransactionsApi.Lottery)                                  // sysTransactions
		sysTransactionsRouterWithoutAuth.POST("checkwin", sysTransactionsApi.CheckWin)                                // sysTransactions
	}
}
