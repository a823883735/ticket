package router

import (
	"ticket/service"
)

func RegisterUserRouter() {
	adminRouter := ginEngine.Group("/v1/user")
	{
		adminRouter.POST("login", service.UserLogin)
		adminRouter.POST("register", service.UserRegister)
		adminRouter.POST("newPwd", service.UserModifyPwd)
	}
}

