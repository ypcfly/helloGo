package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"hello/controllers"
	"hello/models"
	_ "hello/routers"
	"reflect"
)

// 自定义标签
type CustomStruct struct {
	Name     string `customTag:"userName" tags:"user_name"`
	Password string `customTag:"pwd"`
	Age      int    `customTag:"userAge"`
}

func main() {

	cust := CustomStruct{"张三", "123456", 22}
	t := reflect.TypeOf(cust)
	val := reflect.ValueOf(cust)
	for i := 0; i < val.NumField(); i++ {
		println(">>>> field type=" + val.Field(i).Kind().String())
	}

	field, _ := t.FieldByName("Name")
	println(">>>> custom struct field Tag=" + field.Tag + " <<<<")

	println(">>>> struct field name's customTag value=" + field.Tag.Get("customTag") + " <<<<")
	v, _ := field.Tag.Lookup("tags")
	println(">>>> struct field name's tags value=" + v + " <<<<")
	field, _ = t.FieldByName("Password")
	println(">>>> struct field password's tag=" + field.Tag + " <<<<")
	field, _ = t.FieldByName("Age")
	println(">>>> struct field age's tag=" + field.Tag + " <<<<")

	// 过滤器
	beego.InsertFilter("/user/*", beego.BeforeExec, controllers.BeforeExecFilter, false, false)
	beego.InsertFilter("/user/*", beego.BeforeRouter, controllers.BeforeRouterFilter, false, false)
	beego.InsertFilter("/user/*", beego.BeforeStatic, controllers.BeforeStaticFilter, false, false)
	beego.InsertFilter("/user/*", beego.AfterExec, controllers.AfterExecFilter, false, false)
	beego.InsertFilter("/user/*", beego.FinishRouter, controllers.FinishRouterFilter, false, false)

	// run方法前注册错误handler
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}

func init() {
	// 数据库初始化
	models.Init()

	// 初始化session配置
	models.InitSession()
}
