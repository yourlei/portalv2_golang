package router

import (
	"portal/controller"
	"portal/controller/captcha"
	"portal/controller/user"
	"portal/controller/role"
	"portal/controller/app"
	"portal/controller/menu"
	"portal/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.Use(middleware.Cors)
  // 根路由
	router.GET("/", controller.Home)
	// User module
	// 获取验证码
	router.GET("/api/v1/image/base64", captcha.CreatePngData)
	// 用户登录
	router.POST("/api/v1/users/signin", user.Signin)
	// 用户注册
	router.POST("/api/v1/users/signup", user.Signup)
	// 查询用户列表
	router.GET("/api/v1/users", user.QueryUserList)
	// 用户状态变更(启用,禁用,注销)
	router.PATCH("/api/v1/users/status/:id", user.UpdateUserStatus)
	// 审核用户
	router.PATCH("/api/v1/users/check/:id", user.ReviewUser)
	// 编辑用户
	router.PATCH("/api/v1/users/edit/:id", user.EditUser)
	// 更新密码
	router.PATCH("/api/v1/users/password/:id", middleware.SigninRequired, user.ChangePasswd)
	// Role module
	// 角色列表
	router.GET("/api/v1/roles", role.QueryRoleList)
	// 创建角色
	router.POST("/api/v1/roles", role.CreateRole)
	// 编辑角色
	router.PATCH("/api/v1/roles/:id", role.UpdateRole)
	// 删除角色
	router.DELETE("/api/v1/roles/:id", role.DeleteRole)
	// 获取角色组下用户
	router.GET("/api/v1/roles/users/:id", role.GetUserByRole)
	// 迁移用户
	router.POST("/api/v1/roles/users", role.MigrateUser)
	// App module
	// 创建app
	router.POST("/api/v1/app", app.CreateApp)
	// 应用列表
	router.GET("/api/v1/app", app.GetAppList)
	// Menu module
	// 新增菜单
	router.POST("/api/v1/resource/menus", menu.CreateRouter)
	router.DELETE("/api/v1/resource/menus/:id", menu.DeleteRouter)
	// 
	// router.GET("/api/v1/test", controller.Test)
	// listent 3000
	router.Run(":3000")
}