package routers

import (
	"github.com/beego"
	"github.com/nhatvhm/asolution/controllers"
)

func init() {
	beego.Router("/", &ctl.UsersController{}, "get:Index")
	beego.Router("/login", &ctl.LoginController{}, "get,post:Login")
	beego.Router("/logout", &ctl.LoginController{}, "get:Logout")
	beego.Router("/signup", &ctl.LoginController{}, "get,post:Signup")
}
