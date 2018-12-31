package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"hello/models"
	"log"
	"strconv"
)

type UserController struct {
	beego.Controller
}

//配置注解路由
func (u *UserController) URLMapping() {
	u.Mapping("QueryById",u.QueryById)
	u.Mapping("QueryList",u.QueryList)
	u.Mapping("Register",u.Register)
	u.Mapping("SaveUser",u.SaveUser)
	u.Mapping("UserJSON",u.UserJSON)
}
// @router /user [get]
func (c *UserController) Register() {
	logs.Info(">>>> forward to register page start <<<<")


	c.TplName = "user/form.html"
}

// @router /user/get/:id [get]
func (u *UserController) QueryById() {
	strId := u.Ctx.Input.Param(":id")
	logs.Info(">>>> query user by userId start <<<<")
	id,err := strconv.Atoi(strId)
	user := models.QueryUserById(id)
	checkError(err)

	u.Data["User"] = user
	bytes,err := json.Marshal(user)
	u.Ctx.ResponseWriter.Write(bytes)

	//u.TplName = "user/user.html"
}
// @router /user/list [get]
func (u *UserController) QueryList() {
	logs.Info(">>>> query user list start <<<<")

	// 数据库查询所有用户
	users := models.QueryUserList()

	//bytes,err := json.Marshal(users)
	//checkError(err)
	//u.Ctx.ResponseWriter.Write(bytes)

	u.Data["List"] = users
	u.TplName = "user/list.html"
}

// @router /user/save [post]
func (u *UserController) SaveUser() {
	logs.Info(">>>> save register information start <<<<")
	// 获取表单数据
	//form := u.Ctx.Request.PostForm
	//username := form.Get("username")
	//age := form.Get("age")
	//sex := form.Get("sex")
	//mobile := form.Get("mobile")

	// 表单转struct
	var user models.User
	err := u.ParseForm(&user)
	checkError(err)
	// 校验...

	// 写入数据库，返回id值
	id := models.InsertUser(&user)
	println(id)
	//var r orm.RawSeter
	//insert := "insert into t_user (username,age,sex,mobile) values(?,?,?,?)"
	//r = o.Raw(insert,username,age,sex,mobile)
	//_,err := r.Exec()
	//checkError(err)
	u.Ctx.Redirect(302,"/success")
}

/*接收requestBody参数*/
// @router /json/save [post]
func (u *UserController) UserJSON() {
	logs.Info(">>>> save user json information start <<<<")
	// requestBody数据
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody,&user)
	checkError(err)

	// 数据库处理
	id := models.InsertUser(&user)
	println("insert user id=" + strconv.FormatInt(id,10))

	u.Ctx.Redirect(302,"/success")
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}