package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ticket/db"
	"ticket/model"
	"time"
)

// @Summary 新增订单
// @Tags 订单
// @Description add ticket order
// @Accept  multipart/form-data
// @Produce application/json
// @Param  uid	formData    string     true        "用户名"
// @Param  tid	formData    string     true        "门票id"
// @Param  num	formData    string     true        "门票数量"
// @Success 200 {string} string	"ok"
// @Router /v1/order/addOrder [post]
// @contact.name 7947
// @contact.email 7947@qq.com
func AddOrder(ctx *gin.Context) {
	var mod = model.Order{
		UID: ctx.PostForm("uid"),
		TID: ctx.PostForm("tid"),
		Date: time.Now().Format("2000-01-01"),
	}
	var num, err = strconv.ParseInt(ctx.PostForm("num"), 10, 32)
	if err != nil {
		ctx.JSON(ReturnDataFormatErr())
		return
	}
	mod.TNum = int(num)
	ctx.JSON(OperationResult(db.AddOrder(mod)))
}

// @Summary 所有订单列表
// @Tags 订单
// @Description get all order
// @Accept  multipart/form-data
// @Produce application/json
// @Param   index  query    string     true        "页码"
// @Param   size  query    string     true        "页长"
// @Success 200 {string} string	"ok"
// @Router /v1/order/orderList [get]
// @contact.name 7947
// @contact.email 7947@qq.com
func AllOrderList(ctx *gin.Context) {
	var index, err = strconv.ParseInt(ctx.Query("index"), 10, 32)
	if err == nil {
		var size, err = strconv.ParseInt(ctx.Query("size"), 10, 32)
		if err == nil {
			var total, data, err = db.GetMyOrderList(int(index), int(size), "")
			ctx.JSON(ReturnDataList(int(index), int(size), total, data, err))
			return
		}
	}
	ctx.JSON(ReturnDataFormatErr())
}

// @Summary 我的订单列表
// @Tags 订单
// @Description get my order list
// @Accept  multipart/form-data
// @Produce application/json
// @Param   index  query    int     true        "页码"
// @Param   size  query    int     true        "页长"
// @Param   uid  query    string     true        "用户id"
// @Success 200 {string} string	"ok"
// @Router /v1/order/myOrderList [get]
// @contact.name 7947
// @contact.email 7947@qq.com
func MyOrderList(ctx *gin.Context) {
	var uid = ctx.Query("uid")
	if uid == "" {
		ctx.JSON(ReturnNoLogin())
		return
	}
	var index, err = strconv.ParseInt(ctx.Query("index"), 10, 32)
	if err == nil {
		var size, err = strconv.ParseInt(ctx.Query("size"), 10, 32)
		if err == nil {
			var total, data, err = db.GetMyOrderList(int(index), int(size), uid)
			ctx.JSON(ReturnDataList(int(index), int(size), total, data, err))
			return
		}
	}
	ctx.JSON(ReturnDataFormatErr())
}

// @Summary 删除订单
// @Tags 订单
// @Description delete my order
// @Accept  multipart/form-data
// @Produce application/json
// @Param   u_id  query    string     true        "用户id"
// @Param   id  query    int     true        "订单id"
// @Success 200 {string} string	"ok"
// @Router /v1/order/deleteOrderById [delete]
// @contact.name 7947
// @contact.email 7947@qq.com
func DeleteOrderById(ctx *gin.Context) {
	var id, err = strconv.ParseInt(ctx.Query("id"), 10, 32)
	if err != nil {
		ctx.JSON(ReturnDataFormatErr())
		return
	}
	var mod = model.Order{
		Id: int(id),
		UID: ctx.Query("u_id"),
		IsDel: true,
	}
	ctx.JSON(OperationResult(db.DeleteOrder(mod)))
}