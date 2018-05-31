package service

import (
	"log"

	"portal/database"
	"portal/common"
)
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
		AND deleted_at = "0000-01-01 00:00:00"`, email)
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
func Signup(User common.SignupForm)  {
	
}