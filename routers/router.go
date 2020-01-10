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
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/article/:slug([\\w\\-]+).html", &controllers.MainController{}, "get:ArticleInfo")

	// admin
	//beego.InsertFilter("/v1/admin/*", beego.BeforeRouter, FilterUser)
	beego.Router("/v1/admin/login", &admin.UserController{}, "post:Login")
	beego.Router("/v1/admin/getInfo", &admin.UserController{}, "get:GetInfo")
	beego.Router("/v1/admin/logout", &admin.UserController{}, "post:Logout")
	// ----------- tag --------------
	beego.Router("/v1/admin/article/tagList", &admin.ArticleController{}, "get:TagList")
	beego.Router("/v1/admin/article/createTag", &admin.ArticleController{}, "put:CreateTag")
	beego.Router("/v1/admin/article/updateTag", &admin.ArticleController{}, "post:UpdateTag")
	// ----------- category --------------
	beego.Router("/v1/admin/article/categoryList", &admin.ArticleController{}, "get:CategoryList")
	beego.Router("/v1/admin/article/createCategory", &admin.ArticleController{}, "put:CreateCategory")
	beego.Router("/v1/admin/article/updateCategory", &admin.ArticleController{}, "post:UpdateCategory")

	// ----------- article --------------
	beego.Router("/v1/admin/article/articleList", &admin.ArticleController{}, "get:ArticleList")
	beego.Router("/v1/admin/article/getCategories", &admin.ArticleController{}, "get:GetCategories")
	beego.Router("/v1/admin/article/getTags", &admin.ArticleController{}, "get:GetTags")
	beego.Router("/v1/admin/article/detail", &admin.ArticleController{}, "get:Detail")
	beego.Router("/v1/admin/article/createArticle", &admin.ArticleController{}, "put:CreateArticle")
	beego.Router("/v1/admin/article/uploadImage", &admin.ArticleController{}, "post:UploadImage")
	beego.Router("/v1/admin/article/updateArticle", &admin.ArticleController{}, "post:UpdateArticle")

	// ----------- page -----------------
	beego.Router("/v1/admin/page/pageList", &admin.PageController{}, "get:PageList")
	beego.Router("/v1/admin/page/createPage", &admin.PageController{}, "put:CreatePage")
	beego.Router("/v1/admin/page/detail", &admin.PageController{}, "get:Detail")
	beego.Router("/v1/admin/page/updatePage", &admin.PageController{}, "post:UpdatePage")
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
