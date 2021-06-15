package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OperationResult(ok bool, err error) (int, gin.H) {
	if err != nil {
		return http.StatusOK, gin.H{
			"code": 500,
			"msg": "服务器异常",
		}
	}
	if ok {
		return http.StatusOK, gin.H{
			"code": 200,
			"data": ok,
		}
	} else {
		return http.StatusOK, gin.H{
			"code": 400,
			"data": false,
		}
	}
}