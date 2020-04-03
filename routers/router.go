package routers

import (
	"beegoWeb/controllers"
	"beegoWeb/controllers/component"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home/console", &controllers.HomeController{}, "get:Console")
	beego.Router("/home/homepage1", &controllers.HomeController{}, "get:Homepage1")
	beego.Router("/home/homepage2", &controllers.HomeController{}, "get:Homepage2")
	beego.Router("/component/grid/list", &component.GridController{}, "get:List")
}
