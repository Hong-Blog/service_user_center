package sysRole

import (
	"errors"
	"log"
	"usercenter/db"
	"usercenter/models/sysUser"
)

func GetAllRole(request GetAllRoleRequest) (list []SysRole, count int) {
	dataSql := `
select id, name, description, available, create_time, update_time
from sys_role
`
	countSql := "select count(1) from sys_role "
	var params = make([]interface{}, 0)
	var filter string
	if len(request.KeyWord) > 0 {
		filter = " where name like ? "
		params = append(params, "%"+request.KeyWord+"%")
	}
	dataSql += filter + " limit ?,? "
	countSql += filter

	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count sys_role err: ", err.Error())
	}

	offset, limit := request.GetLimit()
	params = append(params, offset)
	params = append(params, limit)
	err = db.Db.Select(&list, dataSql, params...)
	if err != nil {
		log.Panicln("select sys_role err: ", err.Error())
	}

	return
}

func GetAllRoleWithCheckedByUserId(userId int) (list []RoleWithChecked) {
	user := sysUser.SysUser{Id: userId}
	exist := user.ExistById()

	if !exist {
		return make([]RoleWithChecked, 0)
	}

	dataSql := `
select r.*, if(isnull(sur.id), 0, 1) checked
from sys_role r
         left join sys_user_role sur on r.id = sur.role_id and sur.user_id = ?
order by r.id
`
	err := db.Db.Select(&list, dataSql, userId)
	if err != nil {
		log.Panicln("GetAllRoleWithCheckedByUserId err: ", err.Error())
	}
	return
}

func AddRole(request AddRoleRequest) (err error) {
	role := SysRole{}
	role.Name.String = request.Name
	if role.ExistByName() {
		return errors.New("角色已存在")
	}

	insertSql := `
insert into sys_role (name, description, available, create_time, update_time)
values (?, ?, ?, now(), now());
`
	result := db.Db.MustExec(insertSql, request.Name, request.Description, request.Available)
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("添加失败")
	}
	return
}

func (r *SysRole) ExistByName() bool {
	dataSql := `
select ifnull((select 1 from sys_role where name = ? limit 1), 0);
`
	count := 0
	if err := db.Db.QueryRow(dataSql, r.Name.String).Scan(&count); err != nil {
		log.Panicln("sysrole ExistByName err: ", err.Error())
	}
	return count > 0
}

func GetById(id int) (role SysRole) {
	dataSql := `
select id, name, description, available, create_time, update_time
from sys_role
where id = ?
`
	err := db.Db.Get(&role, dataSql, id)
	if err != nil {
		log.Panicln("get role by id err: ", err.Error())
	}
	return
}

func UpdateRole(req UpdateRoleRequest) error {
	dataSql := `
update sys_role
set name        = ?,
    description = ?,
    available   = ?
where id = ?;
`
	result := db.Db.MustExec(dataSql, req.Name, req.Description, req.Available, req.Id)
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func (r *SysRole) DeleteById() error {
	countSql := `
select count(1)
from sys_user_role
where role_id = ?;
`
	count := 0
	if err := db.Db.QueryRow(countSql, r.Id).Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("已分配用户，不可删除！")
	}

	deleteSql := `
delete
from sys_role
where id = ?
`
	result := db.Db.MustExec(deleteSql, r.Id)
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("删除失败")
	}
	return nil
}
