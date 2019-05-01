package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)

//表的设计
type User struct{
	Id int
	Name string
	Pwd string
}

// 文章结构体，不建议字段中间使用大写字母或者下划线
type Article struct {
	Id int
	Aname string
	Atime time.Time
	Acount int
	Acontent string
	Aimg string
}

//


func init()  {
	// 记录数据库基本信息
	orm.RegisterDataBase("default","mysql","root:swl19960706@tcp(58.87.104.82:3306)/test3?charset=utf8")
	// 映射model数据
	orm.RegisterModel(new(User))
	// 生成表
	orm.RunSyncdb("default",false,true)
}


