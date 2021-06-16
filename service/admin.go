package service

import (
	"github.com/gin-gonic/gin"
	"ticket/db"
	"ticket/model"
)

// @Summary 管理员登录
// @Tags 管理员
// @Description get admin login
// @Accept  multipart/form-data
// @Produce application/json
// @Param   userName  formData    string     true        "用户名"
// @Param   password  formData    string     true        "密码"
// @Success 200 {string} string	"ok"
// @Router /v1/admin/login [post]
// @contact.name 7947
// @contact.email 7947@qq.com
func AdminLogin(ctx *gin.Context) {
	var admin model.Admin = model.Admin{
		Id: ctx.PostForm("userName"),
		Pwd: ctx.PostForm("password"),
	}
	ctx.JSON(OperationResult(db.AdminLogin(admin)))
}