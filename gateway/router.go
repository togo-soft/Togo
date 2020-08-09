package gateway

import (
	"Togo/gateway/handler"
	"Togo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Router() *gin.Engine {
	router := gin.Default()
	//开启允许跨域
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//通用路由
	var general = router.Group("/api/general")
	{
		//用户注册功能
		general.POST("/signup", handler.Signup)
		//用户登录功能
		general.POST("/signin", handler.Signin)
	}
	//用户路由 这部分路由需要JWT认证 使用JWT中间件
	var user = router.Group("/api/user", middleware.JWTAuth())
	{
		//查找用户的个人信息
		user.GET("/profile", handler.Profile)
		//用户帐号注销功能
		user.GET("/cancellation", handler.Cancellation)
		//用户退出登录功能
		user.GET("/logout", handler.Cancellation)
		//用户信息更新功能
		user.POST("/update", handler.ModifyInformation)
	}
	//管理员路由
	var admin = router.Group("/api/admin")
	{
		//查找用户列表
		admin.GET("/userlist", handler.FindMany)
	}
	return router
}
