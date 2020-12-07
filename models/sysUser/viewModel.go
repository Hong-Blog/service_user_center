package sysUser

import (
	"usercenter/models"
)

type GetAllUserRequest struct {
	models.PagedRequest
	KeyWord  string `json:"key_word"`
	Username string `json:"username"`
}

type UpdateUserRequest struct {
	Id       int    `uri:"id"`
	Nickname string `json:"nickname"` // 昵称
	Mobile   string `json:"mobile"`   // 手机号
	Email    string `json:"email"`    // 邮箱地址
	Qq       string `json:"qq"`       // QQ
}

type AddUserRequest struct {
	Username string `json:"username" binding:"required" display:"用户名"`
	Password string `json:"password" binding:"required" display:"密码"` // 登录密码
	Nickname string `json:"nickname" binding:"required" display:"昵称"` // 昵称
	Mobile   string `json:"mobile"`                                   // 手机号
	Email    string `json:"email" binding:"omitempty,email"`          // 邮箱地址
	Qq       string `json:"qq"`                                       // QQ
	RegIp    string `json:"reg_ip"`                                   // 注册IP
}

type UpdatePasswordByIdRequest struct {
	Id               int    `json:"id"`
	Username         string `json:"username"`
	Password         string `json:"password"`           // 当前密码
	NewPassword      string `json:"new_password"`       // 新密码
	NewPasswordAgain string `json:"new_password_again"` // 新密码
}

type UpdateUserRoleRequest struct {
	Id     int `json:"id"`
	RoleId int `json:"role_id"`
}
