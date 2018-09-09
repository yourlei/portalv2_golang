package database

import (
	"log"
	"time"
	"database/sql"

	"portal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)
// 定义数据库访问实例
var db *sql.DB
var RedisPool *redis.Pool
// Connect mysql
// dbConfig: "user:password@tcp(127.0.0.1:3306)/dbname"
func initMysql() {
	_db, err := sql.Open("mysql", config.MysqlConfig.URL)
	if err != nil {
		log.Fatal(err)
	}
	db = _db
}
// Close mysql connect
func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
// return conn object
func ConnDB() *sql.DB {
	return db
}
// Connect redis
func initRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConfig.URL, redis.DialPassword(config.RedisConfig.Password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
// 初始化连接
func init() {
	initMysql()
	initRedis()
}