package service

import "github.com/gin-gonic/gin"

// @Summary 新增订单
// @Tags 订单
// @Description add ticket order
// @Accept  multipart/form-data
// @Produce application/json
// @Param   userName  formData    string     true        "用户名"
// @Param   password  formData    string     true        "密码"
// @Success 200 {string} string	"ok"
// @Router /v1/order/addOrder [post]
// @contact.name 7947
// @contact.email 7947@qq.com
func AddOrder(ctx *gin.Context) {

}

// @Summary 所有订单列表
// @Tags 订单
// @Description get admin login
// @Accept  multipart/form-data
// @Produce application/json
// @Param   index  query    string     true        "页码"
// @Param   size  query    string     true        "页长"
// @Success 200 {string} string	"ok"
// @Router /v1/order/orderList [get]
// @contact.name 7947
// @contact.email 7947@qq.com
func AllOrderList(ctx *gin.Context) {

}

// @Summary 我的订单列表
// @Tags 订单
// @Description get admin login
// @Accept  multipart/form-data
// @Produce application/json
// @Param   index  query    string     true        "页码"
// @Param   size  query    string     true        "页长"
// @Success 200 {string} string	"ok"
// @Router /v1/order/myOrderList [get]
// @contact.name 7947
// @contact.email 7947@qq.com
func MyOrderList(ctx *gin.Context) {

}