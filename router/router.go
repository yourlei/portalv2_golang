package router

import (
	"portal/controller"
	"portal/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.Use(middleware.Cors)
  // 根路由
	router.GET("/", controller.Home)
	router.GET("/api/v1/image/base64", controller.Create)
	router.POST("/api/v1/users/signin", controller.Signin)
	// listent 3000
	router.Run(":3000")
}