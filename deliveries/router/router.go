package router

import (
	"Togo/deliveries/handler"
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

	//用户路由 增删改查
	var user = router.Group("/api/v1/user")
	{
		user.POST("/create", handler.CreateUser)
		user.GET("/delete", handler.DeleteUser)
		user.POST("/update", handler.UpdateUser)
		user.GET("/info", handler.UserInfo)
	}
	return router
}
