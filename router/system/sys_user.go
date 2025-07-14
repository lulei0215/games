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
		userRouter.POST("admin_register", baseApi.Register)               //
		userRouter.POST("changePassword", baseApi.ChangePassword)         //
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     //
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               //
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                //
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                //
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) //
		userRouter.POST("resetPassword", baseApi.ResetPassword)           //
		userRouter.PUT("setSelfSetting", baseApi.SetSelfSetting)          //
		//
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)
		userRouterPub.POST("dashboard", baseApi.Dashboard)                             //
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)                //
		userRouterPub.POST("resetWithdrawPassword", baseApi.ResetWithdrawPassword)     //
		userRouterPub.POST("changeWithdrawPassword", baseApi.ChangeWithdrawPassword)   //
		userRouterPub.POST("verify-withdraw-password", baseApi.VerifyWithdrawPassword) //
		userRouterPub.POST("set-withdraw-password", baseApi.SetWithdrawPassword)       //
		userRouterPub.POST("login", baseApi.ApiLogin)                                  //
		userRouterPub.POST("register", baseApi.ApiRegister)                            //
		userRouterPub.POST("sendcode", baseApi.SendCode)                               //
		userRouterPub.POST("bindemail", baseApi.BindeMail)                             //
		userRouterPub.POST("decrypt", baseApi.Decrypt)                                 //
		userRouterPub.POST("encrypt", baseApi.Encrypt)                                 //
		userRouterPub.POST("info", baseApi.Info)                                       //
		userRouterPub.POST("list", baseApi.RobotList)                                  //
		userRouterPub.GET("info", baseApi.GetInfo)                                     //
		userRouterPub.GET("autoLogin", baseApi.AutoLogin)                              //
		userRouterPub.POST("updLang", baseApi.UpdateLang)                              //
		userRouterPub.POST("updAudio", baseApi.UpdateAudio)                            //
		userRouterPub.POST("updateRedisUserData", baseApi.UpdateRedisUserDataSafe)     //
	}
}
