package database

import (
	"portal/model"
)
// create row
func CreateLog(log model.Log) (error) {
	Sql := "INSERT INTO portal_log(`user_id`, `url`, `method`, `host`, `proto`, `ua`, `create_at`) VALUES(?,?,?,?,?,?,?)"
	tx, err := ConnDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(Sql, log.UserId, log.Url, log.Method, log.Host, log.Proto, log.Agent, log.CreateAt)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}