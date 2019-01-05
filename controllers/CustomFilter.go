package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func BeforeExecFilter(ctx *context.Context) {
	logs.Info(">>>> BeforeExec filter start <<<<")

}

func BeforeRouterFilter(ctx *context.Context) {
	logs.Info(">>>> BeforeRouter filter start <<<<")

}

// 第一个执行
func BeforeStaticFilter(ctx *context.Context) {
	logs.Info(">>>> BeforeStatic filter start <<<<")
	ctx.Input.Session("name")
}

func AfterExecFilter(ctx *context.Context) {
	logs.Info(">>>> AfterExec filter start <<<<")

}

func FinishRouterFilter(ctx *context.Context) {
	logs.Info(">>>> FinishRouter filter start <<<<")

}
