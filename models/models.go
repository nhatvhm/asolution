package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id           int       `orm:"column(id);auto"`
	FirstName    string    `orm:"column(firstname);size(50)"`
	LastName     string    `orm:"column(lastname);size(50)"`
	Username     string    `orm:"column(username);size(50)"`
	Email        string    `orm:"column(email);size(255)"`
	Password     string    `orm:"column(password);size(128)"`
	Created_date time.Time `orm:"column(created_date);type(timestamp);auto_now_add"`
	Salt         string    `orm:"size(10)"` // when it comes to password security always use salt!
}

func init() {
	orm.RegisterModel(new(User))
}
