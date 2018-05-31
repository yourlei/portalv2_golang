package main

import (
	"portal/database"
	"portal/router"
)

func main()  {
	database.OpenDB("portal:D024Ad41d8cd98f00b204@tcp(192.168.80.243:3306)/portal")
	router.Run()
}