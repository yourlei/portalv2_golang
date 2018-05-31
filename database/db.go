package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
// 定义数据库访问实例
var db *sql.DB
var err error
// 访问mysql
// dbConfig: "user:password@tcp(127.0.0.1:3306)/dbname"
func OpenDB(dbConfig string) {
	db, err = sql.Open("mysql", dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}

// 关闭数据库DB
func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// 获取数据库DB的连接
func ConnDB() *sql.DB {
	return db
}