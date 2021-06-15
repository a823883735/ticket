package router

import "github.com/gin-gonic/gin"

var ginEngine *gin.Engine = gin.Default()

func Run() {
	ginEngine.Run(":8080")
}

func Server() *gin.Engine {
	return ginEngine
}