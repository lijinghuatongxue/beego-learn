package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)

//表的设计
type User struct{
	Id int
	// 后面添加unique等设置的话，必须删除掉表重新建，设置才会生效
	Name string `"orm:unique"`
	Pwd string
}

// 文章结构体，不建议字段中间使用大写字母或者下划线
type Article struct {
	// `pk` 主键
	// `auto` 自增
	//数据库字段的属性设置
	//设置主键pk
	//自增auto
	//全局唯一 unique
	//设置允许为空 null
	//设置长度  size()
	//设置默认值 default()
	//浮点数精度  digits(12);decimal(4)      12345678.1234
	//时间的设置  date   datetime   time.time    type(date)/type(datetime)
	//每次model更新时，都会对时间自动更新	auto_now
	//只有第一次保存时，才更新时间			    auto_now_add
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(30)"`
	Atime time.Time
	Acount  int `orm:"default(0);null"`
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


