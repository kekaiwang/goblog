package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "wangkekai.com"
	c.Data["Email"] = "wkekai@163.com"
	c.TplName = "index.html"
}
