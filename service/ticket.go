package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ticket/db"
	"ticket/model"
)

// @Summary 新增门票
// @Tags 门票
// @Description add ticket
// @Accept  multipart/form-data
// @Produce application/json
// @Param   title  formData    string     true        "门票名称"
// @Param   details  formData    string     true        "门票详情"
// @Param   price  formData    string     true        "门票价格"
// @Success 200 {string} string	"ok"
// @Router /v1/ticket/addTicket [post]
// @contact.name 7947
// @contact.email 7947@qq.com
func AddTicket(ctx *gin.Context) {
	var mod = model.Ticket{
		Title: ctx.PostForm("title"),
		Details: ctx.PostForm("details"),
	}
	var price, err = strconv.ParseFloat(ctx.PostForm("price"), 10)
	if err != nil {
		ctx.JSON(ReturnDataFormatErr())
		return
	}
	mod.Price = float32(price)
	ctx.JSON(OperationResult(db.AddTicket(mod)))
}

// @Summary 门票列表
// @Tags 门票
// @Description get ticket list
// @Accept  multipart/form-data
// @Produce application/json
// @Success 200 {string} string	"ok"
// @Router /v1/ticket/ticketList [get]
// @contact.name 7947
// @contact.email 7947@qq.com
func TicketList(ctx *gin.Context) {
	ctx.JSON(ReturnDataSet(db.GetTicket()))
}

// @Summary 删除门票类型
// @Tags 门票
// @Description get ticket list
// @Accept  multipart/form-data
// @Produce application/json
// @Param   id  query    int     true        "门票id"
// @Success 200 {string} string	"ok"
// @Router /v1/ticket/delTicketList [delete]
// @contact.name 7947
// @contact.email 7947@qq.com
func DelTicket(ctx *gin.Context) {
	var ticketId, err = strconv.ParseInt(ctx.Query("id"), 10, 32)
	if err != nil {
		ctx.JSON(ReturnDataFormatErr())
		return
	}
	ctx.JSON(OperationResult(db.DeleteTicket(int(ticketId))))
}