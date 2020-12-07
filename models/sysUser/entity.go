package sysUser

import "github.com/guregu/null"

type SysUser struct {
	Id            int         `json:"id"`
	Username      null.String `json:"username" swaggertype:"string"`
	Password      null.String `json:"password" swaggertype:"string"`                             // 登录密码
	Nickname      null.String `json:"nickname" swaggertype:"string"`                             // 昵称
	Mobile        null.String `json:"mobile" swaggertype:"string"`                               // 手机号
	Email         null.String `json:"email" swaggertype:"string"`                                // 邮箱地址
	Qq            null.String `json:"qq" swaggertype:"string"`                                   // QQ
	Birthday      null.Time   `json:"birthday" swaggertype:"string"`                             // 生日
	Gender        null.Int    `json:"gender" swaggertype:"integer"`                              // 性别
	Avatar        null.String `json:"avatar" swaggertype:"string"`                               // 头像地址
	UserType      null.String `json:"user_type" db:"user_type" swaggertype:"string"`             // 超级管理员、管理员、普通用户
	Company       null.String `json:"company" swaggertype:"string"`                              // 公司
	Blog          null.String `json:"blog" swaggertype:"string"`                                 // 个人博客地址
	Location      null.String `json:"location" swaggertype:"string"`                             // 地址
	Source        null.String `json:"source" swaggertype:"string"`                               // 用户来源
	Uuid          null.String `json:"uuid" swaggertype:"string"`                                 // 用户唯一表示(第三方网站)
	Privacy       null.Int    `json:"privacy" swaggertype:"integer"`                             // 隐私（1：公开，0：不公开）
	Notification  null.Int    `json:"notification" swaggertype:"integer"`                        // 通知：(1：通知显示消息详情，2：通知不显示详情)
	Score         null.Int    `json:"score" swaggertype:"integer"`                               // 金币值
	Experience    null.Int    `json:"experience" swaggertype:"integer"`                          // 经验值
	RegIp         null.String `json:"reg_ip" db:"reg_ip" swaggertype:"string"`                   // 注册IP
	LastLoginIp   null.String `json:"last_login_ip" db:"last_login_ip" swaggertype:"string"`     // 最近登录IP
	LastLoginTime null.Time   `json:"last_login_time" db:"last_login_time" swaggertype:"string"` // 最近登录时间
	LoginCount    null.Int    `json:"login_count" db:"login_count" swaggertype:"integer"`        // 登录次数
	Remark        null.String `json:"remark" swaggertype:"string"`                               // 用户备注
	Status        null.Int    `json:"status" swaggertype:"integer"`                              // 用户状态
	CreateTime    null.Time   `json:"create_time" db:"create_time" swaggertype:"string"`         // 注册时间
	UpdateTime    null.Time   `json:"update_time" db:"update_time" swaggertype:"string"`         // 更新时间
}
