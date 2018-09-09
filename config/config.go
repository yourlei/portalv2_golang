package config

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"portal/util"
)

// App basic config
type appConfig struct {
	TokenSecrect string
	TokenMaxAge  int
	AesKey       string
}
// DB connect config 
type dbConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	URL      string
}
// Redis connect config
type redisConfig struct {
	Host      string
	Port      int
	Password  string
	URL       string
	MaxIdle   int
	MaxActive int
}
// init variable
var (
	AppConfig    appConfig
	MysqlConfig  dbConfig
	RedisConfig  redisConfig
)
// 解析json配置文件
var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}
// init database params
func initMysql() {
	util.SetStructByJSON(&MysqlConfig, jsonData["mysql"].(map[string]interface{}))
	
	MysqlConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", MysqlConfig.Username, MysqlConfig.Password, 
	MysqlConfig.Host, MysqlConfig.Port,MysqlConfig.Database)
}
// init app
func initApp() {
	util.SetStructByJSON(&AppConfig, jsonData["app"].(map[string]interface{}))
}
// init redis config
func initRedis() {
	util.SetStructByJSON(&RedisConfig, jsonData["redis"].(map[string]interface{}))
	RedisConfig.URL = fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port)
	
	fmt.Print("redis url: ", RedisConfig.URL)
}
func init() {
	initJSON()
	initMysql()
	initApp()
	initRedis()
}