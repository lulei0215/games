package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MonitorTransfersRouter struct{}

// InitMonitorTransfersRouter  monitorTransfers表
func (s *MonitorTransfersRouter) InitMonitorTransfersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	monitorTransfersRouter := Router.Group("monitorTransfers").Use(middleware.OperationRecord())
	monitorTransfersRouterWithoutRecord := Router.Group("monitorTransfers")
	monitorTransfersRouterWithoutAuth := PublicRouter.Group("monitorTransfers")
	{
		monitorTransfersRouter.POST("createMonitorTransfers", monitorTransfersApi.CreateMonitorTransfers)             // monitorTransfers表
		monitorTransfersRouter.DELETE("deleteMonitorTransfers", monitorTransfersApi.DeleteMonitorTransfers)           // monitorTransfers表
		monitorTransfersRouter.DELETE("deleteMonitorTransfersByIds", monitorTransfersApi.DeleteMonitorTransfersByIds) // monitorTransfers表
		monitorTransfersRouter.PUT("updateMonitorTransfers", monitorTransfersApi.UpdateMonitorTransfers)              // monitorTransfers表
	}
	{
		monitorTransfersRouterWithoutRecord.GET("findMonitorTransfers", monitorTransfersApi.FindMonitorTransfers)       // IDmonitorTransfers表
		monitorTransfersRouterWithoutRecord.GET("getMonitorTransfersList", monitorTransfersApi.GetMonitorTransfersList) // monitorTransfers表
	}
	{
		monitorTransfersRouterWithoutAuth.GET("getMonitorTransfersPublic", monitorTransfersApi.GetMonitorTransfersPublic) // monitorTransfers表
		monitorTransfersRouterWithoutAuth.POST("recharge", monitorTransfersApi.Recharge)                                  // sysTransactions
		monitorTransfersRouterWithoutAuth.POST("transfer", monitorTransfersApi.Transfer)                                  // sysTransactions
	}
}
