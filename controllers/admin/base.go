package admin

import (
	"fmt"
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

func (base *baseController) GetClientIp() string {
	ip := base.Ctx.Input.IP()

	return fmt.Sprintf("%s", ip)
}
