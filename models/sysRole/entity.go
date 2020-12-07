package sysRole

import "github.com/guregu/null"

type SysRole struct {
	Id          int64       `json:"id"`
	Name        null.String `json:"name" swaggertype:"string"` // 角色名
	Description null.String `json:"description" swaggertype:"string"`
	Available   null.Int    `json:"available" swaggertype:"integer"`
	CreateTime  null.Time   `json:"create_time" db:"create_time" swaggertype:"string"` // 添加时间
	UpdateTime  null.Time   `json:"update_time" db:"update_time" swaggertype:"string"` // 更新时间
}
