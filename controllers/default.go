package controllers

import (
	"class/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

 func (c *MainController) Get() {
/* 	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"  */
/*	//有orm对象
	o := orm.NewOrm()
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
// 1  有orm对象
	o := orm.NewOrm()
// 2 查询的对象
	user := models.User{}
// 3 指定查询对象字段
//	user.Id = 1
	user.Name = "233"

// 4 查询
	err := o.Read(&user,"Name")
	 	if err != nil {
	 	logs.Info( "查询失败" ,err)
		 return
	 }
	logs.Info( "查询成功",user)

	c.Data["data"] = "home页"
	c.TplName = "test.html"
} 

func (c *MainController) Post() {
/* 	 	c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"  */
		c.Data["data"] = "abc页面"
		c.TplName = "test.html"
	}
