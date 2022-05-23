package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int `orm:"pk;auto"`
	UserName   string
	Password   string
	Age        int
	Gender     int
	Phone      int64
	Addr       string
	CreateTime time.Time
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User))
}
