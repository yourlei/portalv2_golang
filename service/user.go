package service

import (
	"fmt"
	"log"

	"portal/database"
	"portal/common"

	"golang.org/x/crypto/bcrypt"
)
const deleted_at = "0000-01-01 00:00:00"
// 用户登录
// 验证用户信息,生成token
func Signin(email, passwd string) (int, string) {
	var (
		id int
		name string
		password string
		status int
		checkStatus int
	)
	// 根据email查询用户
	rows, err := database.ConnDB().Query(`SELECT 
		id, name, password, status, check_status 
		FROM users 
		WHERE email = ? 
		AND deleted_at = ?`, email, common.DeletedAt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 遍历结果rows
	for rows.Next() {
		err := rows.Scan(&id, &name, &password, &status, &checkStatus)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, password, status, checkStatus)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return id, name
}

// 用户注册
func Signup(User common.SignupForm) int  {
	log.Print(User)
	var count int
	row, err := database.ConnDB().Query(`
		SELECT COUNT(1) AS count FROM users WHERE (email = ? OR mobile = ?) AND deleted_at = ?
	`, User.Email, User.Mobile, common.DeletedAt)
	// print err
	if err != nil {
		log.Fatal(err)
	}
	// 获取count值
	for row.Next() {
		err := row.Scan(&count)
		if err != nil {
			// fmt.Println(err)
			log.Fatal(err)
		}
	}
	// 邮箱或手机号已注册
	if count > 0 {
		return 1
	}
	passwordOK := "admin"
	// passwordERR := "adminxx"

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
			fmt.Println(err)
	}
	fmt.Println(string(hash), "=+++++++++++++")

	// encodePW := string(hash)  // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	// fmt.Println(encodePW)

	// // 正确密码验证
	// err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
	// if err != nil {
	// 		fmt.Println("pw wrong")
	// } else {
	// 		fmt.Println("pw ok")
	// }

	// // 错误密码验证
	// err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR))
	// if err != nil {
	// 		fmt.Println("pw wrong")
	// } else {
	// 		fmt.Println("pw ok")
	// }

	return 0
}