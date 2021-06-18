package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"ticket/docs"
)

func RegisterSwagRouter() {
	docs.SwaggerInfo.Version = "v1.0.0"
	docs.SwaggerInfo.Title = "在线购票"
	docs.SwaggerInfo.Host = "106.13.237.16:10000"
	docs.SwaggerInfo.BasePath = ""

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}