package api

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	SysTransactionsRouter
	MonitorTransfersRouter
	PaymentTransactionsRouter
	PaymentCallbacksRouter
	UserWithdrawalAccountsRouter
}

var (
	sysTransactionsApi        = api.ApiGroupApp.ApiApiGroup.SysTransactionsApi
	monitorTransfersApi       = api.ApiGroupApp.ApiApiGroup.MonitorTransfersApi
	paymentTransactionsApi    = api.ApiGroupApp.ApiApiGroup.PaymentTransactionsApi
	paymentCallbacksApi       = api.ApiGroupApp.ApiApiGroup.PaymentCallbacksApi
	userWithdrawalAccountsApi = api.ApiGroupApp.ApiApiGroup.UserWithdrawalAccountsApi
)
