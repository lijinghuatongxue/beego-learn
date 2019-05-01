package routers

import (
	"class/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	// 注意：当实现了自定义的get请求方法，请求不会访问默认方法
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{},"get:ShowIndex")
}
