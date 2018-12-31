package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	/*固定路由*/
    //beego.Router("/", &controllers.MainController{})
    //beego.Router("/user",&controllers.UserController{})

    beego.BConfig.EnableGzip = true
    beego.BConfig.RouterCaseSensitive = true
	beego.BConfig.MaxMemory = 1 << 26

	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.CopyRequestBody = true

	/*设置模板取值方式*/
	//beego.BConfig.WebConfig.TemplateLeft = "${"
	//beego.BConfig.WebConfig.TemplateRight = "}"
	/*页面文件路径*/
	//beego.BConfig.WebConfig.ViewsPath="views/user"

	/*注解路由*/
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.CommonController{})
	beego.Include(&controllers.FileController{})
	beego.Include(&controllers.LoginController{})

	/*同时注册多个路由*/
	beego.Include(&controllers.UserController{},&controllers.CommonController{},&controllers.FileController{})
	//beego.ErrorController(&controllers.ErrorController{})
}
