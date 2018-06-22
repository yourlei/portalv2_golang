package service

import (
	"portal/model"
	"portal/database"
)
// Grant privilege
func Grant(arg model.RolePrivilege) error {
	return database.BindRoleRes(arg)
}