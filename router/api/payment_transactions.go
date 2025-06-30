package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaymentTransactionsRouter struct{}

// InitPaymentTransactionsRouter  paymentTransactions
func (s *PaymentTransactionsRouter) InitPaymentTransactionsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	paymentTransactionsRouter := Router.Group("paymentTransactions").Use(middleware.OperationRecord())
	paymentTransactionsRouterWithoutRecord := Router.Group("paymentTransactions")
	paymentTransactionsRouterWithoutAuth := PublicRouter.Group("paymentTransactions")
	{
		paymentTransactionsRouter.POST("createPaymentTransactions", paymentTransactionsApi.CreatePaymentTransactions)             // paymentTransactions
		paymentTransactionsRouter.DELETE("deletePaymentTransactions", paymentTransactionsApi.DeletePaymentTransactions)           // paymentTransactions
		paymentTransactionsRouter.DELETE("deletePaymentTransactionsByIds", paymentTransactionsApi.DeletePaymentTransactionsByIds) // paymentTransactions
		paymentTransactionsRouter.PUT("updatePaymentTransactions", paymentTransactionsApi.UpdatePaymentTransactions)              // paymentTransactions
	}
	{
		paymentTransactionsRouterWithoutRecord.GET("findPaymentTransactions", paymentTransactionsApi.FindPaymentTransactions)       // IDpaymentTransactions
		paymentTransactionsRouterWithoutRecord.GET("getPaymentTransactionsList", paymentTransactionsApi.GetPaymentTransactionsList) // paymentTransactions
	}
	{
		paymentTransactionsRouterWithoutAuth.GET("getPaymentTransactionsPublic", paymentTransactionsApi.GetPaymentTransactionsPublic) // paymentTransactions
		paymentTransactionsRouterWithoutAuth.POST("createTrade", paymentTransactionsApi.CreateTrade)                                  // paymentTransactions
		// paymentTransactionsRouterWithoutAuth.POST("queryTrade", paymentTransactionsApi.QueryTrade)  // paymentTransactions
		paymentTransactionsRouterWithoutAuth.POST("createPayment", paymentTransactionsApi.CreatePayment) // paymentTransactions
		// paymentTransactionsRouterWithoutAuth.POST("queryPayment", paymentTransactionsApi.QueryPayment)  // paymentTransactions
	}
}
