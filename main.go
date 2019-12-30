package main

import (
	"github.com/astaxie/beego"
	_ "github.com/wkekai/goblog/routers"
)

func main() {
	beego.Run()
}

