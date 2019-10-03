package handler

import (
	"Togo/models"
	"Togo/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//用户用例
var UserCS = usecases.NewUser()

//查看用户信息
func UserInfo(this *gin.Context) {
	var user = &models.User{}
	fmt.Println("当前查询的id:", this.Query("id"))
	id, _ := strconv.Atoi(this.Query("id"))
	user = UserCS.FetchUser(id)
	this.JSON(200, gin.H{
		"message": user,
	})
}

//创建用户
func CreateUser(this *gin.Context) {
	var user = &models.User{}
	err := this.Bind(user)
	fmt.Println("获取输入信息", user)
	if err != nil {
		this.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	if UserCS.CreateUser(user) {
		this.JSON(200, gin.H{
			"message": "创建成功",
		})
	} else {
		this.JSON(200, gin.H{
			"message": "创建失败",
		})
	}
}

//删除用户
func DeleteUser(this *gin.Context) {
	fmt.Println("删除的id:", this.Query("id"))
	id, _ := strconv.Atoi(this.Query("id"))
	if err := UserCS.DeleteUser(id); err == nil {
		this.JSON(200, gin.H{
			"message": "删除成功",
		})
	} else {
		this.JSON(200, gin.H{
			"message": err,
		})
	}
}

//修改用户
func UpdateUser(this *gin.Context) {
	var user = &models.User{}
	err := this.Bind(user)
	fmt.Println("获取需要修改的信息", user)
	if err != nil || user.ID == 0 {
		this.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}
	if err := UserCS.UpdateUser(user); err != nil {
		this.JSON(200, gin.H{
			"message": err,
		})
	} else {
		this.JSON(200, gin.H{
			"message": "修改成功",
		})
	}
}
