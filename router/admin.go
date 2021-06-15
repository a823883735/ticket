package router

import (
	"ticket/service"
)

func RegisterAdminRouter() {
	adminRouter := ginEngine.Group("/v1/admin")
	{
		adminRouter.POST("login", service.AdminLogin)
	}
}
