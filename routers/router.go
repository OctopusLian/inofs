package routers

import (
	"inofs/controllers/login"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &login.LoginController{})
}
