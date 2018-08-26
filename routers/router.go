package routers

import (
	"github.com/TimothyZhang023/beeserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
