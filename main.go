package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wkekai/goblog/models"
	_ "github.com/wkekai/goblog/routers"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}

