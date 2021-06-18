package service

import (
	"github.com/gin-gonic/gin"
	"ticket/db"
	"ticket/model"
)

// @Summary 用户登录
// @Tags 用户
// @Description user login
// @Accept  multipart/form-data
// @Produce application/json
// @Param   userName  formData    string     true        "用户名"
// @Param   password  formData    string     true        "密码"
// @Success 200 {string} string	"ok"
// @Router /v1/user/login [post]
// @contact.name 7947
// @contact.email 7947@qq.com
func UserLogin(ctx *gin.Context) {
	var mod model.User = model.User{
		Id: ctx.PostForm("userName"),
		Pwd: ctx.PostForm("password"),
	}
	ctx.JSON(OperationResult(db.UserLogin(mod)))
}

// @Summary 用户注册
// @Tags 用户
// @Description user register
// @Accept  multipart/form-data
// @Produce application/json
// @Param   userName  formData    string     true        "用户名"
// @Param   password  formData    string     true        "密码"
// @Success 200 {string} string	"ok"
// @Router /v1/user/register [post]
// @contact.name 7947
// @contact.email 7947@qq.com
func UserRegister(ctx *gin.Context) {
	var mod = model.User{
		Id: ctx.PostForm("userName"),
		Pwd: ctx.PostForm("password"),
	}
	ctx.JSON(OperationResult(db.UserRegister(mod)))
}

// @Summary 用户修改密码
// @Tags 用户
// @Description modify user password
// @Accept  multipart/form-data
// @Produce application/json
// @Param   userName  formData    string     true        "用户名"
// @Param   oldPwd  formData    string     true        "旧密码"
// @Param   newPwd  formData    string     true        "新密码"
// @Success 200 {string} string	"ok"
// @Router /v1/user/newPwd [put]
// @contact.name 7947
// @contact.email 7947@qq.com
func UserModifyPwd(ctx *gin.Context) {
	var mod = model.User{
		Id: ctx.PostForm("userName"),
		Pwd: ctx.PostForm("newPwd"),
	}
	var oldPwd = ctx.PostForm("oldPwd")
	ctx.JSON(OperationResult(db.UserModifyPwd(mod, oldPwd)))
}