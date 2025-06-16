package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterPub := RouterPub.Group("user").Use()
	userRouterWithoutRecord := Router.Group("user")
	{
		userRouter.POST("admin_register", baseApi.Register)               // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
		userRouter.PUT("setSelfSetting", baseApi.SetSelfSetting)          // 用户界面配置
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)             // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)              // 获取自身信息
		userRouterPub.POST("resetWithdrawPassword", baseApi.ResetWithdrawPassword)   // 重置提现密码
		userRouterPub.POST("changeWithdrawPassword", baseApi.ChangeWithdrawPassword) // 重置提现密码
		userRouterPub.POST("login", baseApi.ApiLogin)                                // 获取自身信息
		userRouterPub.POST("register", baseApi.ApiRegister)                          // 获取自身信息
		userRouterPub.POST("sendcode", baseApi.SendCode)                             // 获取自身信息
		userRouterPub.POST("bindemail", baseApi.BindeMail)                           // 获取自身信息
	}
}
