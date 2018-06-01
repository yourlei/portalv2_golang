// 公用常量
package common

// 参数绑定示例
// form: content-type = x-www-form-urlencoded
// json: content-type = application/json
// binding: required 表示该参数必选且不为空
/* ***************************
 * type User struct {
 *     Username string `form:"username" json:"username" binding:"required"`
 *     Passwd   string `form:"passwd" json:"passwd" binding:"required"`
 *     Age      int    `form:"age" json:"age"`
 * }
 **************************** */
// 用户注册提交表单信息
type SignupForm struct {
	Name      string `json:"name" binding:"required"`
	Email     string `valid:"email~请输入正确的邮箱" json:"email" binding:"required"`
	Mobile    string `valid:"numeric~请输入正确的手机号码" json:"mobile" binding:"required"`
	Password  string `json:"password" binding:"required"`
	RoleId    int    `json:"roleId" binding:"required"`
}
// 登录时提交的请求体
type LoginForm struct {
	Email 	 string `valid:"email~请输入正确的邮箱" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
	Code     string `json:"code" binding:"required"`
}
type ErrorMsg struct {
	msg string
}
// response
type CodeMsg struct {
	code int
	error ErrorMsg
}

// deleted_at 默认值
const DeletedAt = "0000-01-01 00:00:00"