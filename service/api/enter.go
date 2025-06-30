package api

type ServiceGroup struct {
	SysTransactionsService
	MonitorTransfersService
	PaymentTransactionsService
	PaymentCallbacksService
	UserWithdrawalAccountsService
}
