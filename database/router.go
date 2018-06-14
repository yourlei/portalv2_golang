package database

import (
	"fmt"
	"strconv"
	"encoding/json"
	"database/sql"
	"portal/model"
)
var createRouter = "INSERT INTO `portal_router`(`name`,`route`,`type`,`parent`,`priority`,`schema`,`remark`) VALUES(?,?,?,?,?,?,?)"
// Create Router
func CreateRouter(r model.Route) (int, error) {
	var v []byte
	tx, err := ConnDB().Begin()
	if err != nil {
		return 1, err
	}
	// interface to json string
	if v, err = json.Marshal(r.Schema); err != nil {
		return 1, err
	}
	res, err := tx.Exec(createRouter, r.Name, r.Route, r.Type, r.Parent, r.Priority, r.Schema, string(v))
	if err != nil {
		fmt.Println(err)
		return 1, err
	}
	tx.Commit()
	id, _ := res.LastInsertId()
	return int(id), nil
}
// Check router uniqueness
// Router as url params, parent as father menu
// First level menu parent default value -1
func UniqueRouter(r model.Route) (int, error) {
	var (
		Sql = `SELECT name FROM portal_router WHERE `
		name string
		values []string
	)
	_parent := strconv.Itoa(r.Parent)
	// Fisrt level menu
	if r.Parent == -1 {
		Sql += `(route = ? OR name = ?) AND parent = -1`
		values = append(values, r.Route, r.Name)
	} else {
		// chidl menu
		Sql += `(route = ? OR name = ?) AND parent = ?`
		values = append(values, r.Route, r.Name, _parent)
	}
	// slice to interface
	params := make([]interface{}, len(values))
	for i, v := range values {
		params[i] = v
	}
	err := ConnDB().QueryRow(Sql, params...).Scan(&name)
	// IF Not Found
	if err == sql.ErrNoRows {
		return -1, nil
	}
	// error
	if err != nil {
		return 1, err
	}
	return 0, nil
}