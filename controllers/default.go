package controllers

import (
	"class/models"

	//	"class/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	//	"github.com/astaxie/beego/logs"
//	"github.com/astaxie/beego/orm"

	//"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

 func (c *MainController) Get() {
/* 	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"  */

// +++++++++++++++++++++++++++++++++++++++++++++++++  orm insert 操作

	//有orm对象
/*	o := orm.NewOrm()
	// 有一个要插入数据的结构体对象
	user := models.User{}
	// 对结构体对象赋值
	user.Name = "245"
	user.Pwd = "278"

	// 插入
	_,err := o.Insert(&user)
	if err != nil{
		logs.Info("插入失败",err)
		return
		}*/

// +++++++++++++++++++++++++++++++++++++++++++++++++  orm select 操作

// 1.  有orm对象
/*	o := orm.NewOrm()
// 2. 查询的对象
	user := models.User{}
// 3. 指定查询对象字段
//	user.Id = 1
	user.Name = "233"

// 4 查询
	err := o.Read(&user,"Name")
	 	if err != nil {
	 	logs.Info( "查询失败" ,err)
		 return
	 }
	logs.Info( "查询成功",user)*/

// +++++++++++++++++++++++++++++++++++++++++++++++++   orm update 操作
// 1 要有orm对象
/*	o := orm.NewOrm()
// 2 需要更新的结构体对象
	user := models.User{}
// 3 查到需要更新的数据

	user.Id = 1
	err := o.Read(&user)

// 4 给数据重新赋值
	if err == nil {
		user.Name = "121"
		// 如果只更新name的话，就把下面的注释掉，程序依旧可行
		user.Pwd = "333"
	}
// 5 更新
	_,err = o.Update(&user)
	if err != nil{
		logs.Info( "更新失败",err)
		return
	}*/

//+++++++++++++++++++++++++++++++++++++++++++++++++   orm delete 操作

/*// 有orm对象
	o := orm.NewOrm()
// 删除的对象
	user := models.User{}
// 指定删除哪一条数据
	user.Id = 2
// 删除
	_,err := o.Delete(&user)
	if err != nil{
		logs.Info("删除失败",err)
		return
	}
//
*/


//	c.Data["data"] = "home页"
//	c.TplName = "test.html"
    c.TplName = "register.html"
} 
//+++++++++++++++++++++++++++++++++++++++++++++++++ 注册页面
func (c *MainController) Post() {
/* 	 	c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"  */
/*		c.Data["data"] = "abc页面"
		c.TplName = "test.html"*/
// 1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
// 2.对数据进行校验
	if userName == "" || pwd  == "" {
		logs.Info("数据不能为空")
		c.Ctx.WriteString("数据不能为空")
		c.Redirect("/register",302)
		return
	}
// 3.插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_,err := o.Insert(&user)
	if err != nil{
		logs.Info("插入数据库失败")
		c.Redirect("/register",302)
		return
	}
	//logs.Info(userName,pwd)

// 4.返回登陆界面
	c.Ctx.WriteString("注册成功")
	}

// +++++++++++++++++++++++++++++++++++++++++++++++++ 登陆页面
func (c*MainController)ShowLogin(){
	c.TplName = "login.html"
}
func (c*MainController)HandleLogin(){
	//c.Ctx.WriteString("这是登陆界面的post")
	// 1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")

	// 2。 判断数据是否合法
	if userName == ""|| pwd == ""{
		logs.Info("输入不合法")
		c.TplName = "login.html"
		return
	}
	// 3。 查询账号密码是否正确
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	err := o.Read(&user,"Name")
	if err != nil{
		logs.Info("查询数据库失败")

		//指定视图文件，可以传数据
		//c.Data{}
		c.TplName = "login.html"
		//不可以传数据，但是速度快，有状态码
		//c.Redirect("login",302)
		return
	}
	// 4。跳转
	c.Ctx.WriteString("欢迎回来～")
}