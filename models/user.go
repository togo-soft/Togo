package models

import "time"

// 用户表结构
type User struct {
	Id         int64     `form:"-" json:"id"`
	Username   string    `form:"username" json:"username" xorm:"varchar(32) notnull unique"` //用户名
	Nickname   string    `form:"nickname" json:"nickname" xorm:"varchar(32) notnull"`        //昵称
	Password   string    `form:"password" json:"-" xorm:"varchar(64) notnull"`               //密码
	Email      string    `form:"email" json:"email" xorm:"varchar(128) notnull unique"`      //邮箱
	Phone      string    `form:"phone" json:"phone" xorm:"varchar(11)"`                      //电话号码
	CreateTime time.Time `form:"-" json:"create_time" xorm:"created"`                        //注册时间
	LastTime   time.Time `form:"-" json:"last_time" xorm:"updated"`                          //上次登录时间
	LastIp     string    `form:"-" json:"last_ip" xorm:"varchar(46)"`                        //上次登录IP
	Status     bool      `form:"status" json:"status"`                                       //账户状态
}

// 用户组结构
type Group struct {
	Id   int    `form:"-" json:"id" xorm:"pk autoincr"`
	Name string `form:"group_name" json:"group_name" xorm:"varchar(32) notnull"`
	Role string `form:"role" json:"role" xorm:"varchar(32) notnull"`
}
