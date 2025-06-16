package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysTransactionsRouter struct{}

// InitSysTransactionsRouter 初始化 sysTransactions表 路由信息
func (s *SysTransactionsRouter) InitSysTransactionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysTransactionsRouter := Router.Group("sysTransactions").Use(middleware.OperationRecord())
	sysTransactionsRouterWithoutRecord := Router.Group("sysTransactions")
	sysTransactionsRouterWithoutAuth := PublicRouter.Group("sysTransactions")
	{
		sysTransactionsRouter.POST("createSysTransactions", sysTransactionsApi.CreateSysTransactions)             // 新建sysTransactions表
		sysTransactionsRouter.DELETE("deleteSysTransactions", sysTransactionsApi.DeleteSysTransactions)           // 删除sysTransactions表
		sysTransactionsRouter.DELETE("deleteSysTransactionsByIds", sysTransactionsApi.DeleteSysTransactionsByIds) // 批量删除sysTransactions表
		sysTransactionsRouter.PUT("updateSysTransactions", sysTransactionsApi.UpdateSysTransactions)              // 更新sysTransactions表
	}
	{
		sysTransactionsRouterWithoutRecord.GET("findSysTransactions", sysTransactionsApi.FindSysTransactions)       // 根据ID获取sysTransactions表
		sysTransactionsRouterWithoutRecord.GET("getSysTransactionsList", sysTransactionsApi.GetSysTransactionsList) // 获取sysTransactions表列表
	}
	{
		sysTransactionsRouterWithoutAuth.GET("getSysTransactionsPublic", sysTransactionsApi.GetSysTransactionsPublic) // sysTransactions表开放接口
		sysTransactionsRouterWithoutAuth.GET("get", sysTransactionsApi.Get)                                           // sysTransactions表开放接口
	}
}
