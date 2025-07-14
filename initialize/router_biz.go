package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		apiRouter := router.RouterGroupApp.Api
		apiRouter.InitSysTransactionsRouter(privateGroup, publicGroup)
		apiRouter.InitMonitorTransfersRouter(privateGroup, publicGroup)
		apiRouter.InitPaymentTransactionsRouter(privateGroup, publicGroup)
		apiRouter.InitPaymentCallbacksRouter(privateGroup, publicGroup)
		apiRouter.InitUserWithdrawalAccountsRouter(privateGroup, publicGroup)
		apiRouter.InitUserAgentRelationRouter(privateGroup, publicGroup)
		apiRouter.InitGamesRouter(privateGroup, publicGroup)
		apiRouter.InitUserRebatesRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		// 注释掉system路由，避免与api路由冲突
		// {
		// 	systemRouter := router.RouterGroupApp.System
		// 	systemRouter.InitUserAgentRelationRouter(privateGroup, publicGroup)
		// }
		apiRouter.InitUserBetRecordRouter(privateGroup, publicGroup)
	}
}
