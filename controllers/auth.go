package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/nhatvhm/asolution/models"
	"github.com/nhatvhm/asolution/utils"
	"log"
)

func (this *LoginController) LoginView() {
	this.TplName = "authentication/login.html"
}

var sessionName = beego.AppConfig.String("SessionName")

type LoginController struct {
	beego.Controller
}

func (this *LoginController) RegisterView() {
	this.TplName = "authentication/register.html"
}

func (this *LoginController) Register() {
	firstname := this.GetString("firstname")
	lasttname := this.GetString("lastname")
	email := this.GetString("lastname")
	username := this.GetString("username")
	password1 := this.GetString("password1")
	password2 := this.GetString("password2")
	test := models.RegisterForm{Username: username, Password1: password1, Password2: password2, FirstName: firstname, LastName: lasttname, Email: email}

	valid := validation.Validation{}
	b, err := valid.Valid(&test)

	if err != nil {
	}

	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	} else {
		salt := utils.GetRandomString(15)
		encodedPwd := salt + "$" + utils.EncodePassword(password1, salt)

		o := orm.NewOrm()
		o.Using("default")

		user := new(models.User)
		user.Username = username
		user.FirstName = firstname
		user.LastName = lasttname
		user.Email = email
		user.Password = encodedPwd
		user.Salt = salt

		o.Insert(user)

		this.Redirect("/", 302)
	}
	this.TplName = "authentication/register.html"
}

func (this *LoginController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	var user models.User
	if VerifyUser(&user, username, password) {
		v := this.GetSession(sessionName)
		if v == nil {
			this.SetSession(sessionName, user.Id)
		}
		this.Redirect("/app/welcome", 302)

	} else {
		this.Redirect("/register", 302)
	}

}

func (this *LoginController) Logout() {
	this.DelSession(sessionName)
	this.Redirect("/login", 302)
}

func (this *LoginController) AppView() {
	this.TplName = "app/welcome.html"
}

func VerifyUser(user *models.User, username, password string) (success bool) {
	// search user by username or email
	if HasUser(user, username) == false {
		return
	}
	if VerifyPassword(password, user.Password) {
		// success
		success = true
	}
	return
}

func HasUser(user *models.User, username string) bool {
	var err error
	qs := orm.NewOrm()
	user.Username = username
	err = qs.Read(user, "Username")
	if err == nil {
		return true
	}
	return false
}

func VerifyPassword(rawPwd, encodedPwd string) bool {
	// split
	var salt, encoded string
	salt = encodedPwd[:10]
	encoded = encodedPwd[11:]

	return utils.EncodePassword(rawPwd, salt) == encoded
}

// customize filters for fine grain authorization
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session(sessionName).(int)
	if !ok && ctx.Input.Uri() != "/login" && ctx.Input.Uri() != "/register" {
		ctx.Redirect(302, "/login")
	}
}
