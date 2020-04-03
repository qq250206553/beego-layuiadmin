package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Console() {
	c.Layout = "admin/base.html"
	c.TplName = "admin/home/console.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "modules/home/console.html"
}
func (c *HomeController) Homepage1() {
	c.Layout = "admin/base.html"
	c.TplName = "admin/home/homepage1.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "modules/home/homepage1.html"
}
func (c *HomeController) Homepage2() {
	c.Layout = "admin/base.html"
	c.TplName = "admin/home/homepage2.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "modules/home/homepage2.html"
}
