package handler

import (
	"Togo/usecases"
	"github.com/gin-gonic/gin"
)

//uuc user use case
var uuc = usecases.NewUserUC()

//用户-注册
func Signup(this *gin.Context) {
	this.JSON(uuc.Signup(this))
}

//用户-登陆
func Signin(this *gin.Context)  {
	this.JSON(uuc.Signin(this))
}

//用户-退出
func Logout(this *gin.Context)  {
	this.JSON(uuc.Logout(this))
}

//用户-个人中心
func Profile(this *gin.Context)  {
	this.JSON(uuc.Profile(this))
}


//查看用户信息
func FindOne(this *gin.Context) {
	this.JSON(uuc.FindOne(this))
}

//查看用户列表
func FindMany(this *gin.Context) {
	r := uuc.FindMany(this)
	this.JSON(r.Code, r)
}

//删除用户
func Cancellation(this *gin.Context) {
	this.JSON(uuc.Cancellation(this))
}

//修改用户
func ModifyInformation(this *gin.Context) {
	this.JSON(uuc.ModifyInformation(this))
}