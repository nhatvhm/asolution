package routers

import (
	"github.com/astaxie/beego"
	"asolution/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:LoginView;post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/register", &controllers.LoginController{}, "get:RegisterView;post:Register")
	beego.Router("/app/welcome", &controllers.LoginController{}, "get:AppView")

	beego.InsertFilter("/*", beego.BeforeRouter, controllers.FilterUser)

}
