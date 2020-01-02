package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

func Init() {
	dbhost := beego.AppConfig.String("mysqlurls")
	dbport := beego.AppConfig.String("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	dbuser := beego.AppConfig.String("mysqluser")
	dbpassword := beego.AppConfig.String("mysqlpass")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dsn, 30)
	orm.RegisterModelWithPrefix("wkk_", new(AdminUser))
}
