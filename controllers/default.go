package controllers

import (
	"class/models"
	"path"
	_ "path/filepath"
	"time"

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
		//c.Ctx.WriteString("数据不能为空")
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
	c.Redirect("/login",302)
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
	//c.Ctx.WriteString("欢迎回来～")
	c.Redirect("/index",302)
}
// 显示首页内容
func (c*MainController)ShowIndex(){
	o :=orm.NewOrm()
	//构建数组
	var articles []models.Article
	_,err := o.QueryTable("Article").All(&articles)
	if err != nil{
		logs.Info("查询所有文章失败")
		return
	}
	// 打印拿到的是什么
	logs.Info(articles)
	// 传递拿到的数据
	c.Data["articles"] = articles
	c.TplName = "index.html"

}

// 显示添加文章界面
func (c*MainController)ShowAdd(){
	c.TplName = "add.html"
}

// 添加处理文章界面
func (c*MainController)HandleAdd(){
	// 1. 拿到数据
	/*	ArtiName := c.GetString("articleName")
		Acount := c.GetString("content")*/
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	//logs.Info(Acount,artiName)
	f,h,err := c.GetFile("uploadname")
	//f 暂时不用，先关闭
	defer f.Close()
	//+++++++++++++++++++++++++++++++++++++++++++++++++  图片处理
	// 打印图片后缀
	fileext := path.Ext(h.Filename)
	logs.Info(fileext)
	// 图片大小限制
	if h.Size > 5000000000 {
		logs.Info("上传文件过大")
		return
	}
	// 上传文件重命名，防止重复
	// 2006-01-02 15:04:05 是golang语言诞生时间
	filename := time.Now().Format("2006-01-02 15:04:05")+fileext

	if err != nil{
		logs.Info("上传文件失败")
		return
	}else {
		c.SaveToFile("uploadname","./static/img/"+filename)
		//logs.Info(artiName,artiContent)
	}
	//2. 判断数据是否为空
	if artiName == "" || artiContent == ""{
		logs.Info("文章标题不能为空")
		return
	}
	// 3. 插入数据
	o := orm.NewOrm()
	arti := models.Article{}
	arti.ArtiName = artiName
	arti.Acontent = artiContent
	arti.Aimg = "./static/img/"+filename
	_,err = o.Insert(&arti)
	if err != nil {
		logs.Info("文章插入数据库失败")
		return
	}
	// 4.返回文章界面
	c.Redirect("/index",302)
}

// 显示内容详情页
func (c*MainController)ShowContent(){
	// 1.获取详情页
	id,err := c.GetInt("id")
	logs.Info("++++++++++++++++++ Id is ",id)
	if err != nil{
		logs.Info("获取文章id失败",err)
		return
	}
	//2. 数据库查询
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		logs.Info("查询ID错误",err)
		return
	}
	// 3.传递数据给视图
	c.Data["article"] = arti
	c.TplName = "content.html"
}

// 显示编辑界面

func (c*MainController)ShowUpdate(){
	id,err := c.GetInt("id")
	logs.Info("++++++++++++++++++ Id is ",id)
	if err != nil{
		logs.Info("获取文章id失败",err)
		return
	}
	//2. 数据库查询
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		logs.Info("查询ID错误",err)
		return
	}
	// 3.传递数据给视图
	c.Data["article"] = arti
	c.TplName = "update.html"
}
// 处理更新业务代码
func (c*MainController)HandleUpdate(){
	// 1. 拿到数据
	id,_ := c.GetInt("id")
	articleName := c.GetString("articleName")
	content := c.GetString("content")
	// 获取图片
	f,h,err := c.GetFile("uploadname")
	var filename string
	if err != nil{
		logs.Info("上传文件失败")
		return
	}else {
		//f 暂时不用，先关闭
		defer f.Close()
		//+++++++++++++++++++++++++++++++++++++++++++++++++  图片处理
		// 打印图片后缀
		fileext := path.Ext(h.Filename)
		logs.Info(fileext)
		// 图片大小限制
		if h.Size > 5000000000 {
			logs.Info("上传文件过大")
			return
		}
		// 上传文件重命名，防止重复
		// 2006-01-02 15:04:05 是golang语言诞生时间
		filename = time.Now().Format("2006-01-02 15:04:05")+fileext
		c.SaveToFile("uploadname","./static/img/"+filename)
		//logs.Info(artiName,artiContent)
	}
	// 2. 对数据进行处理
	if articleName == "" || content == ""{
		logs.Info("更新数据获取失败")
		return
	}
	// 3. 更新数据
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		logs.Info("查询数据库错误")
	}
	arti.ArtiName = articleName
	arti.Acontent = content
	arti.Aimg = "./static/img/"+filename

	_,err = o.Update(&arti,"ArtiName","Acontent","Aimg")
	if err != nil{
		logs.Info("更新数据库失败")
		return
	}

	// 4. 返回列表数据
	c.Redirect("/index",302)

}
// 删除文章业务
func (c*MainController) HandleDelete(){
	// 拿到数据
	id,err := c.GetInt("id")
	if err != nil{
		logs.Info("获取id数据失败")
		return
	}
	// 执行删除操作
	o := orm.NewOrm()
	arti := models.Article{Id:id}
	err = o.Read(&arti)
	if err != nil{
		logs.Info("查询错误")
		return
	}
	o.Delete(&arti)
	// 返回列表页
	c.Redirect("/index",302)
}

type DataController struct {
	beego.Controller
}

type LIKE struct {
	Title string
	Project []string
	Configure string
	Ssh_port string
	Ip string
	Service []string
	Use_port []string
	Ping string
	Heathy string
	Ssh string
}

//type JSONS struct {
//	//必须的大写开头
//	Code string
//	Msg  string
//	User []string `json:"user_info"`//key重命名,最外面是反引号
//	Like LIKE
//}

//func (c *DataController) Get() {
//	data := &JSONS{
//		"200",
//		"获取成功",
//		[]string{"maple","18"},
//		LIKE{
//			"192.168.0.17",
//		    "2c4g100G标准硬盘",
//		    "鸿蒙"}}
//	c.Data["json"] = data
//	c.ServeJSON()
//}
func (c *DataController) Get() {
	data := &LIKE{
			"esxi",
			[]string{"鸿蒙","女我不甜","夸父"},
			"2c4g100G标准硬盘",
	        "232",
			"192.168.0.23",
			[]string{"tomcat","nginx"},
			[]string{"443","80","8080"},"32","99","22",
	}
	c.Data["json"] = data
	c.ServeJSON()
}