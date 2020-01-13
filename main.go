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
	orm.Debug = true

	beego.AddFuncMap("sub", sub)
	beego.AddFuncMap("add", add)

	beego.Run()
}

func sub(in int)(out int){
	out = in - 1
	return
}

func add(in int)(out int){
	out = in + 1
	return
}

