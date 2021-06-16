package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	ok, err := db.AdminLogin(admin)
	if err != nil {
		log.Println("admin login error: ", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg": "服务器异常",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": ok,
	})
	ctx.JSON(OperationResult(db.AdminLogin(admin)))
}