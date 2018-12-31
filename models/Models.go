package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 数据库
func Init() {
	logs.Info(">>>> database init start <<<<")

	//  数据库驱动
	orm.RegisterDriver("postgres",orm.DRPostgres)
	// 注册数据库
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	url := beego.AppConfig.String("url")
	dbname := beego.AppConfig.String("dbname")
	datasource := "postgres://" + username + ":" + password + "@" + url + "/" + dbname + "?sslmode=disable"
	orm.RegisterDataBase("default","postgres",datasource)
	// 最大连接和空闲连接
	orm.SetMaxOpenConns("default",30)
	orm.SetMaxIdleConns("default",10)
	// 注册model
	orm.RegisterModel(new(User))
}

