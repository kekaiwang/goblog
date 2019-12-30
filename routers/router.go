package routers

import (
	"github.com/astaxie/beego"
	"github.com/wkekai/goblog/controllers"
)

const (
	ONE_DAY = 24 * 3600
)

func init() {
	// session config
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "SESSIONID"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = ONE_DAY
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	//blog
    beego.Router("/", &controllers.MainController{})

	// admin
	//beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
}

//var FilterUser = func(ctx *context.Context) {
//	val, ok := ctx.Input.Session(background.SESSIONNAME).(string)
//
//	if !ok || val == "" {
//		if ctx.Request.Method == "GET" {
//			ctx.Redirect(302, "login")
//		} else if ctx.Request.Method == "POST" {
//			resp := helper.NewResponse()
//			resp.Status = RS.RS_user_not_login
//			resp.Data = "/login"
//			resp.WriteJson(ctx.ResponseWriter)
//		}
//	}
//}
