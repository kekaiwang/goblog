package main

import (
	"github.com/astaxie/beego"
	"github.com/wkekai/goblog/models"
	_ "github.com/wkekai/goblog/routers"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}

