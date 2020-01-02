package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o orm.Ormer
}

func (base *baseController) Prepare() {
	base.o = orm.NewOrm()
}
