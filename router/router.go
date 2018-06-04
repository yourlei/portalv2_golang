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
	// 获取验证码
	router.GET("/api/v1/image/base64", controller.Create)
	// 用户登录
	router.POST("/api/v1/users/signin", controller.Signin)
	// 用户注册
	router.POST("/api/v1/users/signup", controller.Signup)
	// 查询用户列表
	router.GET("/api/v1/users", controller.QueryUser)
	// listent 3000
	router.Run(":3000")
}