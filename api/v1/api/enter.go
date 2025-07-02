package api

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SysTransactionsApi
	MonitorTransfersApi
	PaymentTransactionsApi
	PaymentCallbacksApi
	UserWithdrawalAccountsApi
	UserAgentRelationApi
	GamesApi
}

var (
	sysTransactionsService        = service.ServiceGroupApp.ApiServiceGroup.SysTransactionsService
	monitorTransfersService       = service.ServiceGroupApp.ApiServiceGroup.MonitorTransfersService
	paymentTransactionsService    = service.ServiceGroupApp.ApiServiceGroup.PaymentTransactionsService
	paymentCallbacksService       = service.ServiceGroupApp.ApiServiceGroup.PaymentCallbacksService
	userWithdrawalAccountsService = service.ServiceGroupApp.ApiServiceGroup.UserWithdrawalAccountsService
	userAgentRelationService      = service.ServiceGroupApp.ApiServiceGroup.UserAgentRelationService
	gamesService                  = service.ServiceGroupApp.ApiServiceGroup.GamesService
)
