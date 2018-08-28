package database

import (
	"time"
	
	"portal/util"
	"portal/model"
)
// 新增用户
const addUser = "INSERT INTO" +
								" `portal_user`" +
								"(`name`, `email`, `mobile`, `password`, `created_at`, `updated_at`)" +
								" VALUES(?, ?, ?, ?, ?, ?)"
// Select one user
const findOne = "SELECT `id`, `name`, `email`, `mobile`, `status`, `check_status` FROM portal_user WHERE "
// 插入用户ID和角色ID
const insertUserRole = "INSERT INTO `portal_user_role`(`role_id`, `user_id`) VALUES(?, ?)"
// 查询用户列表
const selectUser = "SELECT"                 +
									" `u`.`id`,"              +
									" `u`.`name`,"            +
									" `u`.`email`,"           +
									" `u`.`mobile`,"          + 
									" `u`.`status`,"          +
									" `u`.`remark`,"          +
									" `u`.`check_status`,"    +
									" `u`.`check_remark`,"    +
									" `r`.`name` AS role,"    +
									" `u`.`created_at`,"      +
									" `u`.`updated_at`"       +
									" FROM portal_user AS u"  +
									" INNER JOIN portal_user_role AS ur ON u.id = ur.user_id" +
									" INNER JOIN portal_role AS r ON r.id = ur.role_id"       +
									" WHERE "
// 用户登录
func Signin(email string) (data model.TokenAndUser, err error) {
	// 根据email查询用户
	Sql := `SELECT id, name, password, status, check_status FROM portal_user WHERE email = ? AND status != 3`
	if err := ConnDB().QueryRow(Sql, email).Scan(
		&data.Id, 
		&data.Name,
		&data.Password,
		// &data.Mobile,
		&data.Status,
		&data.CheckStatus,
	); err != nil {
		return data, err
	}
	// select role id
	if err := ConnDB().QueryRow(`SELECT role_id FROM portal_user_role WHERE user_id = ?`, 
		data.Id).Scan(&data.RoleId); err != nil {
			
	}
	return data, nil
}
// 新增用户
func AddUser(User model.SignupForm) error {
	tx, txErr := ConnDB().Begin()
	if txErr != nil {
		return txErr
	}
	// add row to user table
	result, err1 := tx.Exec(addUser, User.Name, User.Email, User.Mobile, User.Password, time.Now().Format(util.TimeFormat), time.Now().Format(util.TimeFormat))
	if err1 != nil {
		return err1
	}
	// update user_role table
	userId, _ := result.LastInsertId()
	_, err2 := tx.Exec(insertUserRole, User.RoleId, userId)	
	if  err2 != nil {
		return err2
	}
	return tx.Commit()
}
// Check email or mabile had registered
func FindOneUser(where string, query ...interface{}) (model.User, bool) {
	var user model.User
	if err := ConnDB().QueryRow(findOne+where, query...).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Mobile,
		&user.Status,
		&user.CheckStatus,
		); err != nil {
		return user, false
	}
	return user, true
}
// 查询用户列表
func FindAllUser(where string, query ...interface{}) (data []interface{}, err error) {
	var result = make([]interface{}, 0)
	rows, err := ConnDB().Query(selectUser + where, query...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 遍历行, 追加到result slice
	for rows.Next() {
		var	data = &model.User{}
		if err = rows.Scan(
			&data.Id,
			&data.Name,
			&data.Email,
			&data.Mobile,
			&data.Status,
			&data.Remark,
			&data.CheckStatus,
			&data.CheckRemark,
			&data.Role,
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
// 更新用户状态
func UpdateUserStatus(id int, status int, remark string) error {
	var err error
	// IF status = 3 then cancel accoutn, set deleted_at value
	if 3 != status {
		if remark != "" {
			_, err = ConnDB().Exec("UPDATE portal_user SET `status` = ?, `remark` = ? WHERE id = ?", status, remark, id)
		} else {
			_, err = ConnDB().Exec("UPDATE portal_user SET `status` = ? WHERE id = ?", status, id)
		}
	} else {
		_, err = ConnDB().Exec("UPDATE portal_user SET `status` = ?, `remark` = ?, `deleted_at` = ? WHERE id = ?", status, remark, time.Now().Format(util.TimeFormat), id)
	}
	// IF error
	if err != nil {
		return err
	}
	return nil
}
// 审核用户
func ReviewUser(id int, status int, remark string) error {
	var err error
	// with remark
	if remark != "" {
		_, err = ConnDB().Exec("UPDATE portal_user SET `check_status` = ?, `check_remark` = ? WHERE id = ?", status, remark, id)
	} else {
		_, err = ConnDB().Exec("UPDATE portal_user SET `check_status` = ? WHERE id = ?", status, id)
	}
	// IF error
	if err != nil {
		return err
	}
	return nil
}
// 编辑用户
func EditUser(id int, sql string) error {
	stmt, err := ConnDB().Prepare(sql)
	// IF error
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	// IF error
	if err != nil {
		return err
	}
	return nil
}
// change password
func ChangePasswd(id int, passwd string) error{
	stmt, err := ConnDB().Prepare(`UPDATE portal_user SET password = ? WHERE id = ?`)
	// IF error
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(passwd, id)
	// IF error
	if err != nil {
		return err
	}
	return nil
}
// select password
func GetPasswd(id int) (string, error){
	var pw string
	row := ConnDB().QueryRow(`SELECT password FROM portal_user WHERE id = ?`, id)
	if  err := row.Scan(&pw); err != nil {
		return "", err
	}
	return pw, nil
}