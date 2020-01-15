package routers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/wkekai/goblog/RS"
	"github.com/wkekai/goblog/controllers"
	"github.com/wkekai/goblog/controllers/admin"
	"github.com/wkekai/goblog/helper"
	"github.com/wkekai/goblog/models"
	"html/template"
	"net/http"
)

const (
	ONE_DAY = 24 * 3600
)

func init() {
	// session config
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "SESSIONWKK"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = ONE_DAY
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	//blog
    beego.Router("/", &controllers.MainController{})
	beego.Router("/page/:slug([\\w]+).html", &controllers.MainController{}, "get:PageInfo")
	beego.Router("/article/:slug([\\w\\-]+).html", &controllers.MainController{}, "get:ArticleInfo")
	beego.Router("/archives.html", &controllers.MainController{}, "get:Archives")
	beego.Router("/categories/:link([\\w]+).html", &controllers.MainController{}, "get:Categories")
	beego.Router("/tags/:link([\\w]+).html", &controllers.MainController{}, "get:Categories")

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

	beego.ErrorHandler("404", HttpNotFound)
}

func HttpNotFound(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/404.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var FilterUser = func(ctx *context.Context) {
	token := ctx.Input.Cookie("Admin-Token")

	resp := helper.NewResponse()

	if token == "" {
		var user admin.UserParams

		err := json.Unmarshal(ctx.Input.RequestBody, &user)

		if err != nil {
			resp.Status = RS.RS_user_not_login
			resp.Tips(helper.ERROR, "信息错误")
			resp.WriteJson(ctx.ResponseWriter)
			return
		}

		status := models.Login(user.Username, user.Password, ctx.Input.IP())

		resp.Status = status.Status

		switch status.Status {
		case 204:
			resp.Status = RS.RS_user_not_login
			resp.Tips(helper.ERROR, "管理员账号错误")
			resp.WriteJson(ctx.ResponseWriter)
			return
		case 401:
			resp.Status = RS.RS_user_not_login
			resp.Tips(helper.WARNING, "管理员账号/密码错误")
			resp.WriteJson(ctx.ResponseWriter)
			return
		}
	} else {
		if token != beego.AppConfig.String("Token") {
			resp.Status = RS.RS_user_not_login
			resp.Tips(helper.WARNING, "管理员账号/密码错误")
			resp.WriteJson(ctx.ResponseWriter)
			return
		}
	}
}
