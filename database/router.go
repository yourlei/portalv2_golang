package database

import (
	"portal/model"
)
var createRouter = `INSERT INTO portal_router(name,router,type,parent,priority,schema,remark) VALUES(?,?,?,?,?,?,?)`
// Create Router
func CreateRouter(router model.Router) (int, error) {
	_, err := ConnDB().Exec(createRouter, router.Name, router.Router, router.Type, router.Parent, router.Priority, router.Schema, router.Remark)
	// _, err := ConnDB().Exec(createRouter, router...)
	if err != nil {
		return 1, err
	}
	return 0, nil
}