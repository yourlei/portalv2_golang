package test

import (
	"testing"
	// "os"
	// "path/filepath"
	"portal/service"
	// "portal/database"
)

func TestCreateApp(t *testing.T) {
	// database.OpenDB("root:scut2018@tcp(192.168.80.243:3306)/portal2?parseTime=true")
	var name string = "测试_A"
	code, err := service.CreateApp(name)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("success code: ", code)
	}
	// exepath, err := os.Executable()
	// if err != nil {
	// 	t.Error(err)
	// }
	// file := filepath.Dir(exepath)
	
	// t.Log("myfile ", file)
}