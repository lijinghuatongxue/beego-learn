package controllers

import (
	"class/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (this*ArticleController)ShowArticleList(){
	// 有一个orm对象
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles[]models.Article
	qs.All(&articles) //select * from Article
	count ,err:=qs.Count()//返回数据条数
	// 获取总页数
	pageSize := 1
	pageCount := float64(count)/float64(pageSize)

	if err != nil{
		logs.Info("查询错误")
		return
	}
	logs.Info("count=",count)

	this.Data["count"] = count
	this.Data["pageCount"] = pageCount
}
func (this*ArticleController) Get() {
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
	this.TplName = "register.html"
}



//+++++++++++++++++++++++++++++++++++++++++++++++++ 注册页面
func (this*ArticleController) Post() {
	/* 	 	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"  */
	/*		c.Data["data"] = "abc页面"
			c.TplName = "test.html"*/
	// 1.拿到数据
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")
	// 2.对数据进行校验
	if userName == "" || pwd  == "" {
		logs.Info("数据不能为空")
		//c.Ctx.WriteString("数据不能为空")
		this.Redirect("/register",302)
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
		this.Redirect("/register",302)
		return
	}
	//logs.Info(userName,pwd)

	// 4.返回登陆界面
	this.Redirect("/login",302)
	this.Ctx.WriteString("注册成功")
}

// +++++++++++++++++++++++++++++++++++++++++++++++++ 登陆页面
func (this*ArticleController)ShowLogin(){
	this.TplName = "login.html"
}
func (this*ArticleController)HandleLogin(){
	//c.Ctx.WriteString("这是登陆界面的post")
	// 1.拿到数据
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")

	// 2。 判断数据是否合法
	if userName == ""|| pwd == ""{
		logs.Info("输入不合法")
		this.TplName = "login.html"
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
		this.TplName = "login.html"
		//不可以传数据，但是速度快，有状态码
		//c.Redirect("login",302)
		return
	}
	// 4。跳转
	//c.Ctx.WriteString("欢迎回来～")
	this.Redirect("/index",302)
}
// 显示首页内容
func (this*ArticleController)ShowIndex(){
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
	this.Data["articles"] = articles
	this.TplName = "index.html"

}

// 显示添加文章界面
func (this*ArticleController)ShowAdd(){
	this.TplName = "add.html"
}

// 添加处理文章界面
func (this*ArticleController)HandleAdd(){
	// 1. 拿到数据
	/*	ArtiName := c.GetString("articleName")
		Acount := c.GetString("content")*/
	artiName := this.GetString("articleName")
	artiContent := this.GetString("content")
	//logs.Info(Acount,artiName)
	f,h,err := this.GetFile("uploadname")
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
		this.SaveToFile("uploadname","./static/img/"+filename)
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
	this.Redirect("/index",302)
}

// 显示内容详情页
func (this*ArticleController)ShowContent(){
	// 1.获取详情页
	id,err := this.GetInt("id")
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
	this.Data["article"] = arti
	this.TplName = "content.html"
}

// 显示编辑界面

func (this*ArticleController)ShowUpdate(){
	id,err := this.GetInt("id")
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
	this.Data["article"] = arti
	this.TplName = "update.html"
}
// 处理更新业务代码
func (this*ArticleController)HandleUpdate(){
	// 1. 拿到数据
	id,_ := this.GetInt("id")
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	// 获取图片
	f,h,err := this.GetFile("uploadname")
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
		this.SaveToFile("uploadname","./static/img/"+filename)
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
	this.Redirect("/index",302)

}
// 删除文章业务
func (this*ArticleController) HandleDelete(){
	// 拿到数据
	id,err := this.GetInt("id")
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
	this.Redirect("/index",302)
}