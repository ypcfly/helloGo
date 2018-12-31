package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"hello/controllers"
	"hello/models"
	_ "hello/routers"
)


func main() {
	// 数据库初始化
	models.Init()

	// 过滤器
	//beego.InsertFilter("/*",beego.BeforeExec,controllers.UserFilter)

	// 错误handler
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()


}


//func init() {
//	// set default database
//	username := beego.AppConfig.String("username")
//	password := beego.AppConfig.String("password")
//	url := beego.AppConfig.String("url")
//	dbname := beego.AppConfig.String("dbname")
//
//	orm.RegisterDataBase("default", "postgres", "postgres://" + username + ":" + password + "@" + url + "/" + dbname + "?sslmode=disable", 30)
//}


