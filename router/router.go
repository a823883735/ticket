package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket/conf"
)

var ginEngine *gin.Engine = gin.Default()

func Run() {
	ginEngine.Use(Cors())
	RegisterAdminRouter()
	RegisterUserRouter()
	RegisterSwagRouter()
	RegisterTicketRouter()
	RegisterOrderRouter()
	port := conf.Conf().Section("server").Key("port").Value()
	if port == "" {
		port = "80"
	}
	ginEngine.Run(":" + port)
}

func Server() *gin.Engine {
	return ginEngine
}

//	Allow cross-domain processing.
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		//	ignore all methods option function.
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		//	handle request.
		ctx.Next()
	}
}