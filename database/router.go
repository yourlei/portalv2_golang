package database

import (
	"portal/util"
	"time"
	"errors"
	"strconv"
	"encoding/json"
	"database/sql"
	"portal/model"
)
// Create Router
var createRouter = "INSERT INTO `portal_router`(`name`,`route`,`type`,`parent`,`priority`,`schema`,"+
									 " `remark`, `created_at`, `updated_at`) VALUES(?,?,?,?,?,?,?,?,?)"
var selectRouter = "SELECT `id`, `name`, `route`, `parent`, `schema`, `priority`, `type`, `created_at`," +
									 " `updated_at` FROM portal_router WHERE "
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
	type router struct {
	Id        int       `json:"id"`                              //ID
	AppId     string    `json:"appid"`                           //所属应用
	Name      string    `json:"name"`                            //名称
	Route     string    `json:"item"`                            //路由地址
	Type      int       `json:"action"`                          //路由类型
	Parent    int       `json:"parent"`                          //父级id
	Priority  int       `json:"priority"`                        //权重
	Schema    string    `json:"schema"`                          //参数配置
	Remark    string    `json:"remark"`                          //描述
	CreatedAt time.Time `json:"created_at"`                      //创建时间
	UpdatedAt time.Time `json:"updated_at"`                      //更新时间
}
	var result = make([]interface{}, 0)
	rows, err := ConnDB().Query(selectRouter + where, query...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 遍历行, 追加到result slice
	for rows.Next() {
		// var	data = &model.Route{}
		var	data = &router{}
		if err = rows.Scan(
			&data.Id,
			&data.Name,
			&data.Route,
			&data.Parent,
			&data.Schema,
			&data.Priority,
			&data.Type,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			return result, err
		} else {
			result = append(result, data)
		}
	}
	return result, nil
}