package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"hello/controllers"
	"hello/models"
	_ "hello/routers"
)

func main() {

	// 过滤器
	beego.InsertFilter("/*", beego.BeforeExec, controllers.BeforeExecFilter)
	beego.InsertFilter("/*", beego.BeforeRouter, controllers.BeforeRouterFilter)
	beego.InsertFilter("/*", beego.BeforeStatic, controllers.BeforeStaticFilter)
	beego.InsertFilter("/*", beego.AfterExec, controllers.AfterExecFilter)
	beego.InsertFilter("/*", beego.FinishRouter, controllers.FinishRouterFilter)

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
