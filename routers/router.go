package routers

import (
	"github.com/astaxie/beego"
	"github.com/wkekai/goblog/controllers"
	"github.com/wkekai/goblog/controllers/admin"
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
	//beego.InsertFilter("/v1/admin/*", beego.BeforeRouter, FilterUser)
	beego.Router("/v1/admin/login", &admin.UserController{}, "post:Login")
	beego.Router("/v1/admin/getInfo", &admin.UserController{}, "get:GetInfo")
	beego.Router("/v1/admin/logout", &admin.UserController{}, "post:Logout")
	beego.Router("/v1/admin/article/tagList", &admin.ArticleController{}, "get:TagList")
	beego.Router("/v1/admin/article/createTag", &admin.ArticleController{}, "put:CreateTag")
	beego.Router("/v1/admin/article/updateTag", &admin.ArticleController{}, "post:UpdateTag")
	beego.Router("/v1/admin/article/modifyTag", &admin.ArticleController{}, "post:ModifyTag")
	beego.Router("/v1/admin/article/articleList", &admin.ArticleController{}, "get:ArticleList")
}

//var FilterUser = func(ctx *context.Context) {
//	val, ok := ctx.Input.Session(admin.SESSIONNAME).(string)
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
