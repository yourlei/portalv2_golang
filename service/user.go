package service

import (
	"fmt"
	"log"
	"errors"

	"portal/database"
	"portal/common"
	"portal/model"

	"golang.org/x/crypto/bcrypt"
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
		FROM portal_users 
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
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return id, name
}

// 用户注册
// @return (code int, msg error)
func Signup(User common.SignupForm) (int, error)  {
	// 验证邮箱,手机号是否已注册
	if flag, _ := database.FindOneUser(User, "(`email` = ? OR `mobile` = ?) AND `deleted_at` = ?"); flag {
		return 100010, errors.New("该邮箱或是手机号已注册")
	}
	// 密码加密
	hash, _ := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	User.Password = string(hash)

	err := database.AddUser(User)
	if (err != nil) {
		return 1, err
	}
	return 0, nil
}

func QueryUserList() ([]*model.User, error) {
	result, err := database.FindAllUser("`email` = ? OR `id` = ? ", "admin@ibbd.net", 1)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(result[0])
	return result, nil
}