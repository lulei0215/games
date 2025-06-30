package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaymentCallbacksRouter struct{}

// InitPaymentCallbacksRouter  paymentCallbacks
func (s *PaymentCallbacksRouter) InitPaymentCallbacksRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	paymentCallbacksRouter := Router.Group("callback").Use(middleware.OperationRecord())
	paymentCallbacksRouterWithoutRecord := Router.Group("callback")
	paymentCallbacksRouterWithoutAuth := PublicRouter.Group("callback")
	{
		paymentCallbacksRouter.POST("createPaymentCallbacks", paymentCallbacksApi.CreatePaymentCallbacks)             // paymentCallbacks
		paymentCallbacksRouter.DELETE("deletePaymentCallbacks", paymentCallbacksApi.DeletePaymentCallbacks)           // paymentCallbacks
		paymentCallbacksRouter.DELETE("deletePaymentCallbacksByIds", paymentCallbacksApi.DeletePaymentCallbacksByIds) // paymentCallbacks
		paymentCallbacksRouter.PUT("updatePaymentCallbacks", paymentCallbacksApi.UpdatePaymentCallbacks)              // paymentCallbacks
	}
	{
		paymentCallbacksRouterWithoutRecord.GET("findPaymentCallbacks", paymentCallbacksApi.FindPaymentCallbacks)       // IDpaymentCallbacks
		paymentCallbacksRouterWithoutRecord.GET("getPaymentCallbacksList", paymentCallbacksApi.GetPaymentCallbacksList) // paymentCallbacks
	}
	{
		paymentCallbacksRouterWithoutAuth.GET("getPaymentCallbacksPublic", paymentCallbacksApi.GetPaymentCallbacksPublic) // paymentCallbacks
		paymentCallbacksRouterWithoutAuth.POST("trade", paymentCallbacksApi.TradeCallback)                                // paymentCallbacks
		paymentCallbacksRouterWithoutAuth.POST("payment", paymentCallbacksApi.PaymentCallback)                            // paymentCallbacks
	}
}
