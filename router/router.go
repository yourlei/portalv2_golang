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
	router.GET("/api/v1/image/base64", controller.CreatePngData)
	// 用户登录
	router.POST("/api/v1/users/signin", controller.Signin)
	// 用户注册
	router.POST("/api/v1/users/signup", controller.Signup)
	// 查询用户列表
	router.GET("/api/v1/users", controller.QueryUser)
	// 用户状态变更(启用,禁用,注销)
	router.PATCH("/api/v1/users/status/:id", controller.UpdateUserStatus)
	// 审核用户
	router.PATCH("/api/v1/users/check/:id", controller.ReviewUser)
	// 编辑用户
	router.PATCH("/api/v1/users/edit/:id", controller.EditUser)
	// 更新密码
	router.PATCH("/api/v1/users/password/:id", controller.ChangePasswd)
	
	router.GET("/api/v1/test", controller.Test)
	// listent 3000
	router.Run(":3000")
}