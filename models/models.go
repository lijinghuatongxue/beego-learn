package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type User struct{
	Id int
	Name string
	Pwd string
}
func init()  {
	// 记录数据库基本信息
	orm.RegisterDataBase("default","mysql","root:swl19960706@tcp(58.87.104.82:3306)/test3?charset=utf8")
	// 映射model数据
	orm.RegisterModel(new(User))
	// 生成表
	orm.RunSyncdb("default",false,true)
}


