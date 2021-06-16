package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// http返回操作成功
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

// http返回数据列表
func ReturnDataSet(data interface{}, err error) (int, gin.H) {
	if err != nil {
		return http.StatusOK, gin.H{
			"code": 500,
			"msg": "服务器异常",
		}
	}
	return http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	}
}

// http数据格式错误
func ReturnDataFormatErr() (int, gin.H) {
	return http.StatusOK, gin.H{
		"code": 400,
		"msg": "数据格式错误",
	}
}

// 未登录
func ReturnNoLogin() (int, gin.H) {
	return http.StatusOK, gin.H{
		"code": 400,
		"msg": "未登录",
	}
}