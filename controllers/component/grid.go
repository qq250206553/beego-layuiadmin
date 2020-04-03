package component

import (
	"github.com/astaxie/beego"
)

type GridController struct {
	beego.Controller
}

func (c *GridController) List() {
	c.Layout = "admin/base.html"
	c.TplName = "admin/component/grid/list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "modules/component/grid/list.html"
}
