package main

import (
	"fmt"
	"os"
	"log"

	// "portal/config"
	_ "portal/database"
	"portal/router"
)
// Define init work
// Loggin to file
func init() {
	logErr, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE, 0755)
	// logWarn, err := os.OpenFile("warnning.log", os.O_RDWR|os.O_CREATE, 0755)
	// logInfo, err := os.OpenFile("portal.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("%s\t\n", err.Error)
	}
	defer logErr.Close()
	logger := log.New(logErr, "\r\n",log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("hello")
}
func main()  {
	// database.OpenDB(config.MysqlConfig.URL)
	router.Run()
}