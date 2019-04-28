package routers

import (
	"class/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
}
