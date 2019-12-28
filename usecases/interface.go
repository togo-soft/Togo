package usecases

import (
	"github.com/gin-gonic/gin"
)

// UserInterface 是用户模块的接口
type UserInterface interface {
	//用户注册接口
	Signup(ctx *gin.Context) (int, *Response)
	//用户登陆接口
	Signin(ctx *gin.Context) (int, *Response)
	//用户退出接口
	Logout(ctx *gin.Context) (int, *Response)

	//用户个人中心
	Profile(ctx *gin.Context) (int, *Response)
	//用户自定义删除接口
	Cancellation(ctx *gin.Context) (int, *Response)
	//修改用户信息接口
	ModifyInformation(ctx *gin.Context) (int, *Response)
	//查询单一用户接口
	FindOne(ctx *gin.Context) (int, *Response)
	//查询用户列表接口
	FindMany(ctx *gin.Context) *List
}

// FileInterface 是文件模块的接口
type FileInterface interface {
	//上传文件
	Upload()
	//下载文件
	Download()
	//分享文件
	Shared()
	//移动文件
	Move()
	//重命名文件
	Rename()
	//收藏文件
	Collage()
}
