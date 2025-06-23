package api

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SysTransactionsApi
	MonitorTransfersApi
}

var (
	sysTransactionsService  = service.ServiceGroupApp.ApiServiceGroup.SysTransactionsService
	monitorTransfersService = service.ServiceGroupApp.ApiServiceGroup.MonitorTransfersService
)
