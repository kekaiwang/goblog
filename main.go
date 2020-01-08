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
	beego.Run()
}

