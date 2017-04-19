package models

import (
	"github.com/astaxie/beego/validation"
)

type RegisterForm struct {
	Username  string `valid:"AlphaNumeric" valid:Required;MinSize(5);MaxSize(25)"`
	FirstName string `valid:"AlphaNumeric"`
	LastName  string `valid:"AlphaNumeric"`
	Email     string `form:"type(email)"    valid:Required"`
	Password1 string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	Password2 string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
}

func (form *RegisterForm) Valid(v *validation.Validation) {
	if form.Password1 != form.Password2 {
		v.SetError("Password2", "Password fields did not match")
	}
}
