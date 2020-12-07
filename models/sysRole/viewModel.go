package sysRole

import "usercenter/models"

type GetAllRoleRequest struct {
	models.PagedRequest
	KeyWord string `json:"key_word"`
}

type RoleWithChecked struct {
	SysRole
	Checked bool `json:"checked"`
}

type AddRoleRequest struct {
	Name        string `json:"name" swaggertype:"string" binding:"required"`        // 角色名称
	Description string `json:"description" swaggertype:"string" binding:"required"` // 角色描述
	Available   *int   `json:"available" swaggertype:"integer" binding:"required"`  // 是否可用 0 不可用 1 可用
}

type UpdateRoleRequest struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" swaggertype:"string" binding:"required"`        // 角色名称
	Description string `json:"description" swaggertype:"string" binding:"required"` // 角色描述
	Available   *int   `json:"available" swaggertype:"integer" binding:"required"`  // 是否可用 0 不可用 1 可用
}
