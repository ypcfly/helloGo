package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CommonController struct {
	beego.Controller
}

//配置注解路由
func (u *CommonController) URLMapping() {
	u.Mapping("Success",u.Success)
	u.Mapping("False",u.False)
}


// @router /success [get]
func (c *CommonController) Success() {
	logs.Info(">>>> forward to success page start <<<<")
	c.TplName = "success.html"
}

// @router /false [get]
func (c *CommonController) False() {

	logs.Info(">>>> forward to false page start <<<<")
	c.TplName = "false.html"
}