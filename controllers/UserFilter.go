package controllers

import "github.com/astaxie/beego/context"

var UserFilter = func(ctx *context.Context) {
	println(">>>> login filter start <<<<")
	_, ok := ctx.Input.Session("username").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}
