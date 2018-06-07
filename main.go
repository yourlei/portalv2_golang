package main

import (
	"portal/database"
	"portal/router"
)

func main()  {
	database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	router.Run()
}