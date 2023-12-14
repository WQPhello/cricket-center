package model

import "time"

// 用户表
type User struct {
	ID        int        `query:"id" xorm:"autoincr" json:"id"`          //新增用户、修改用户的时候不要填写
	Username  string     `query:"username" xorm:"pk" json:"username"`    //用户名
	Password  string     `query:"password" json:"password"`              //用户密码
	CreatedAt *time.Time `xorm:"created DATETIME(6)" json:"created_at" ` //新增用户、修改用户的时候不要填写
	UpdatedAt *time.Time `xorm:"updated DATETIME(6)" json:"updated_at" ` //新增用户、修改用户的时候不要填写
}
