package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		Title: ctx.PostForm("tile"),
		Details: ctx.PostForm("details"),
	}
	var price, err = strconv.ParseFloat(ctx.PostForm("price"), 10)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": "数据有误",
		})
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

}