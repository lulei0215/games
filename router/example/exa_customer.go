package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   //
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    //
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) //
	}
	{
		customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)         //
		customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) //
	}
}
