package service

import (
	"log"

	"portal/database"
)
// 用户登录
// 验证用户信息,生成token
func Signin() (int, string) {
	log.Print("hi, miss you")

	var (
		id int
		name string
	)
	rows, err := database.ConnDB().Query("select id, name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return id, name
}