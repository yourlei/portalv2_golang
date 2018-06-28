package database

import (
	"strings"
	"time"
	"errors"
	"strconv"
	"encoding/json"
	"database/sql"
	"portal/util"
	"portal/model"
)
// Create Router
var createRouter = "INSERT INTO `portal_router`(`name`,`route`,`type`,`parent`,`priority`,`schema`,"+
									 " `remark`, `created_at`, `updated_at`) VALUES(?,?,?,?,?,?,?,?,?)"
// Select route list
var selectRouter =  "SELECT" +
											" a.id," +
											" a.name," +
											" a.route," +
											" a.parent," +
											" a.`schema`," +
											" a.priority," +
											" a.type," +
											" a.created_at," +
											" a.updated_at," +
											" c.app" +
										" FROM" +
											" portal_router a" +
											" JOIN portal_resource AS b ON a.id = b.resource_id" +
											" JOIN portal_app AS c ON b.app_id = c.uuid" +
											" WHERE a.deleted_at = '" + util.DeletedAt + "'"
// Create Router
func CreateRouter(r model.Route) (int, error) {
	var v []byte
	// check parent id
	if r.Parent != -1 {
		if code, _ := FindById(r.Parent, `portal_router`); code == -1 {
			return 1, errors.New("父级ID不存在")
		}
	}
	tx, err := ConnDB().Begin()
	if err != nil {
		return 1, err
	}
	// interface to json string
	if v, err = json.Marshal(r.Schema); err != nil {
		return 1, err
	}
	res, err := tx.Exec(createRouter, r.Name, r.Route, r.Type, r.Parent, r.Priority, 
		string(v), r.Remark, time.Now().Format(util.TimeFormat), time.Now().Format(util.TimeFormat))
	if err != nil {
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
	// Fisrt level menu
	if r.Parent == -1 {
		Sql += `(route = ? OR name = ?) AND parent = -1`
		values = append(values, r.Route, r.Name)
	} else {
		_parent := strconv.Itoa(r.Parent)
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
// Query router list
// return menu router table row
func FindAllRouter(where string, query ...interface{}) ([]interface{}, error) {
	var result = make([]interface{}, 0)
	rows, err := ConnDB().Query(selectRouter + where, query...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 遍历行, 追加到result slice
	for rows.Next() {
		var	data = &model.Route{}
		if err = rows.Scan(
			&data.Id,
			&data.Name,
			&data.Route,
			&data.Parent,
			&data.SchemaTo,
			&data.Priority,
			&data.Type,
			&data.CreatedAt,
			&data.UpdatedAt,
			&data.AppId,
		); err != nil {
			return result, err
		} else {
			result = append(result, data)
		}
	}
	return result, nil
}
// Check child menu appid equal parent appid
func EqualAppid(parentId int, appid string) (bool, error) {
	var (
		appId string
		Sql = `SELECT app_id FROM portal_resource WHERE resource_id = ?`
	)
	err := ConnDB().QueryRow(Sql, parentId).Scan(&appId)
	// not found or error
	if err != nil {
		return false, err
	}
	if appid != appId {
		return false, nil
	}
	return true, nil
}
// Search all parent menu
func FindParentRouter() ([]interface{}, error) {
	type List struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	var result = make([]interface{}, 0)
	Sql := "SELECT `id`, `name` FROM portal_router WHERE parent = -1 AND deleted_at = '0000-01-01 00:00:00'"
	rows, err := ConnDB().Query(Sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// parse
	for rows.Next() {
		var ele = List{}
		if err := rows.Scan(
			&ele.Id,
			&ele.Name,
		); err != nil {
			return result, err
		} else {
			result = append(result, ele)
		}
	}
	return result, nil
}
// Update menu router
func UpdateRouter(id int, r model.RouteUpdate) (int, interface{}) {
	var (
		name string
		where []string
		Sql = "SELECT name FROM portal_router WHERE id != ?"
	)
	// check name or route
	if r.Name != "" || r.Route != "" {
		Sql += " AND (name = '" + r.Name + "' OR route = '" + r.Route + "')" + " AND parent = " + 	strconv.Itoa(r.Parent)
	  err := ConnDB().QueryRow(Sql, id).Scan(&name)
		// IF error
		if err != nil && err != sql.ErrNoRows  {
			return -1, err
		}
	}
	if name != "" {
		return 1, "名称或地址已占用"
	}
	// update
	if r.Name != "" {
		sql := " `name` = '" + r.Name + "'"
		where = append(where, sql)
	}
	if r.Route != "" {
		sql := " `route` = '" + r.Route + "'"
		where = append(where, sql)
	}
	if r.Priority != 0 {
		sql := " `priority` = '" + strconv.Itoa(r.Priority) + "'"
		where = append(where, sql)
	}
	if r.Schema != nil {
		schema, err := json.Marshal(r.Schema)
		if err != nil {
			return 1, err
		}
		sql := " `schema` = '" + string(schema) + "'"
		where = append(where, sql)
	}
	// Have update field
	if len(where) > 0 {
		updateSql := "UPDATE portal_router SET " + strings.Join(where, ",")
		_, err = ConnDB().Exec(updateSql + " WHERE id = ?", strconv.Itoa(id))
		if err != nil {
			return 1, err
		}
	}
	return 0, nil
}