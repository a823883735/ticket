package router

import (
	"github.com/gin-gonic/gin"
	"ticket/conf"
)

var ginEngine *gin.Engine = gin.Default()

func Run() {
	RegisterAdminRouter()
	RegisterUserRouter()
	RegisterSwagRouter()
	port := conf.Conf().Section("server").Key("port").Value()
	if port == "" {
		port = "80"
	}
	ginEngine.Run(":" + port)
}

func Server() *gin.Engine {
	return ginEngine
}