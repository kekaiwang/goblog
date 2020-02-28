package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wkekai/goblog/models"
	_ "github.com/wkekai/goblog/routers"
)

func init() {
	models.Init()
}

func main() {
	orm.Debug = false

	beego.AddFuncMap("sub", sub)
	beego.AddFuncMap("add", add)

	// logs.Async()
	// logs.SetLogger(logs.AdapterMultiFile, `{"filename": "logs/test.log"}`)
	// logs.SetLogger(logs.AdapterConn, `{"net":"tcp","addr":":7020"}`)

	// if runmode is dev use this
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }

	beego.Run()
}

func sub(in int) (out int) {
	out = in - 1
	return
}

func add(in int) (out int) {
	out = in + 1
	return
}
