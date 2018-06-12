package main

import (
	"portal/config"
	"portal/database"
	"portal/router"
)

func main()  {
	database.OpenDB(config.MysqlConfig.URL)
	router.Run()
}