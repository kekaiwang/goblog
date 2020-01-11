package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/wkekai/goblog/models"
)

type BaseController struct {
	beego.Controller
	o orm.Ormer
}

func (base *BaseController) Prepare() {
	base.o = orm.NewOrm()

	// log reqeust info
	base.DoRequest()
}

func (base *BaseController) DoRequest() {
	request := models.NewRequest(base.Ctx.Input)
	models.RequestM.Ch <- request
}
