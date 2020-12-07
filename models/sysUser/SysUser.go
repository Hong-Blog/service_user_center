package sysUser

import (
	"errors"
	"log"
	"usercenter/db"
	"usercenter/utils"
)

func GetAllUser(request GetAllUserRequest) (list []SysUser, count int) {
	strSql := `
select id,
       username,
       password,
       nickname,
       mobile,
       email,
       qq,
       birthday,
       gender,
       avatar,
       user_type,
       company,
       blog,
       location,
       source,
       uuid,
       privacy,
       notification,
       score,
       experience,
       reg_ip,
       last_login_ip,
       last_login_time,
       login_count,
       remark,
       status,
       create_time,
       update_time
from sys_user
`
	var params = make([]interface{}, 0)
	var filter string
	filter += " where is_deleted = false "
	if len(request.KeyWord) != 0 {
		filter += " and username like ? "
		params = append(params, "%"+request.KeyWord+"%")
	}
	if len(request.Username) != 0 {
		filter += " and username = ? "
		params = append(params, request.Username)
	}
	strSql += filter + " order by id desc limit ?, ?;"
	countSql := "select count(1) from sys_user " + filter
	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count sys_user err: ", err.Error())
	}

	offset, limit := request.GetLimit()
	params = append(params, offset)
	params = append(params, limit)
	err = db.Db.Select(&list, strSql, params...)
	if err != nil {
		log.Panicln("select sys_user err: ", err.Error())
	}

	return
}

func GetById(id int) (user SysUser) {
	sysUser := SysUser{}
	dataSql := `
select id,
       username,
       password,
       nickname,
       mobile,
       email,
       qq,
       birthday,
       gender,
       avatar,
       user_type,
       company,
       blog,
       location,
       source,
       uuid,
       privacy,
       notification,
       score,
       experience,
       reg_ip,
       last_login_ip,
       last_login_time,
       login_count,
       remark,
       status,
       create_time,
       update_time
from sys_user
where id = ?
`
	err := db.Db.Get(&sysUser, dataSql, id)
	if err != nil {
		log.Panicln("get user by id err: ", err.Error())
	}
	return sysUser
}

func UpdateUser(request UpdateUserRequest) (success bool) {
	updateSql := `
update sys_user
set nickname = ?,
    mobile = ?,
    email = ?,
    qq = ?
where id = ?;
`
	result, err := db.Db.Exec(updateSql, request.Nickname, request.Mobile, request.Email, request.Qq, request.Id)
	if err != nil {
		log.Panicln("update user by id err: ", err.Error())
	}
	affected, err1 := result.RowsAffected()
	if err1 != nil {
		log.Panicln("not support affected err: ", err1.Error())
	}
	return affected > 0
}

func AddUser(request AddUserRequest) (success bool, err error) {
	exsitUser := SysUser{}
	exsitUser.Username.String = request.Username
	if exsitUser.ExistByName() {
		err = errors.New("用户已经存在")
		return false, err
	}

	insertSql := `
INSERT INTO sys_user (
  username,
  PASSWORD,
  nickname,
  mobile,
  email,
  qq,
  reg_ip,
  user_type,
  STATUS,
  create_time
)
VALUES
  (?, ?, ?, ?, ?, ?, ?,'ADMIN', 1, NOW());
`
	encryptedPassword := utils.PasswordEncrypt(request.Password, request.Username)
	result, err := db.Db.Exec(insertSql, request.Username, encryptedPassword, request.Nickname,
		request.Mobile, request.Email, request.Qq, request.RegIp)
	if err != nil {
		log.Panicln("add user  err: ", err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Panicln("not support affected err: ", err.Error())
	}
	return affected > 0, nil
}

func (u *SysUser) ExistByName() bool {
	dataSql := `
select ifnull((select 1 from sys_user where username = ? limit 1), 0);
`
	count := 0
	err := db.Db.QueryRow(dataSql, u.Username.String).Scan(&count)
	if err != nil {
		log.Panicln("ExistByName err: ", err.Error())
	}
	return count > 0
}

func (u *SysUser) LogicalDeleteById() (err error) {
	dataSql := `
update sys_user
set is_deleted = true
where id = ?
`
	result, sqlErr := db.Db.Exec(dataSql, u.Id)
	if sqlErr != nil {
		log.Panicln("LogicalDeleteById err: ", sqlErr.Error())
		return sqlErr
	}
	affected, affectedErr := result.RowsAffected()
	if affectedErr != nil {
		log.Panicln("not support affected err: ", affectedErr.Error())
		return affectedErr
	}
	if affected == 0 {
		return errors.New("删除用户失败")
	}
	return nil
}

func UpdatePasswordById(request UpdatePasswordByIdRequest) (err error) {
	updateSql := `
update sys_user
set password = ?
where id = ?;
`
	existSql := `
select ifnull((select 1
               from sys_user
               where id = ?
                 and password = ?
               limit 1), 0)
`
	if request.NewPassword != request.NewPasswordAgain {
		return errors.New("输入的密码不一致")
	}

	passwordEncrypt := utils.PasswordEncrypt(request.Password, request.Username)
	count := 0
	if err := db.Db.QueryRow(existSql, request.Id, passwordEncrypt).Scan(&count); err != nil {
		log.Panicln("UpdatePasswordById err: ", err.Error())
		return err
	}

	if count == 0 {
		return errors.New("当前密码不正确")
	}

	newPassowrdEncrypt := utils.PasswordEncrypt(request.NewPassword, request.Username)
	result, err := db.Db.Exec(updateSql, newPassowrdEncrypt, request.Id)
	if err != nil {
		log.Panicln("UpdatePasswordById err: ", err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Panicln("not support affected err: ", err.Error())
		return err
	}
	if affected == 0 {
		return errors.New("更新密码失败")
	}
	return nil
}

func (u *SysUser) ExistById() bool {
	exist := false
	checkExistSql := `
select ifnull((select 1
               from sys_user
               where id = ?
               limit 1), 0) exist
`
	if checkErr := db.Db.QueryRow(checkExistSql, u.Id).Scan(&exist); checkErr != nil {
		log.Panicln("GetAllRoleWithCheckedByUserId err: ", checkErr.Error())
	}
	return exist
}

func UpdateUserRole(request UpdateUserRoleRequest) (success bool) {
	deleteSql := `
delete
from sys_user_role
where user_id = ?;
`
	updateSql := `
insert into sys_user_role
    (user_id, role_id)
    VALUE (?, ?);
`
	tx := db.Db.MustBegin()
	defer tx.Rollback()
	tx.MustExec(deleteSql, request.Id)
	tx.MustExec(updateSql, request.Id, request.RoleId)
	err := tx.Commit()
	if err != nil {
		log.Panicln("UpdateUserRole err: ", err.Error())
	}

	return true
}
