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
// db config 
type dbConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	URL      string
}
// init variable
var (
	AppConfig appConfig
	MysqlConfig  dbConfig
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
	
	// root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", MysqlConfig.Username, MysqlConfig.Password, 
	MysqlConfig.Host, MysqlConfig.Port,MysqlConfig.Database)
	MysqlConfig.URL = url
}
// init app
func initApp() {
	util.SetStructByJSON(&AppConfig, jsonData["app"].(map[string]interface{}))
}

func init() {
	initJSON()
	initMysql()
	initApp()
}