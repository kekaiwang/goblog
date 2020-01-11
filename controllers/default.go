package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/wkekai/goblog/models"
	"strings"
	"time"
)

type MainController struct {
	BaseController
}

type ArticleInfo struct {
	Id int
	Author string
	Title string
	Count int
	Markdown string
	Content string
	CategoryId int
	TagIds string
	Excerpt string
	Previous int
	Next int
	Preview int
	Slug string
	IsDraft int
	Created time.Time
	EditTime time.Time
	Updated time.Time
	DisplayTime time.Time
	CategoryName  string
}

type ArticleList struct {
	Data interface{}
	Total int64
}

func (c *MainController) Get() {
	var articleList ArticleList
	var articleInfos []ArticleInfo

	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select(
		"article.*",
		"category.name as category_name").
		From("wkk_article as article").
		LeftJoin("wkk_category as category").On("article.category_id = category.id").
		Where("article.is_draft < 3").
		OrderBy("article.id").Desc().
		Limit(10).Offset(0).String()

	c.o.Raw(sql).QueryRows(&articleInfos)

	articleList.Total, _ = c.o.QueryTable(new(models.Article)).Filter("is_draft", 2).Count()
	articleList.Data = articleInfos

	c.Data["Data"] = articleList.Data
	c.Data["Title"] = "blog"
	c.Data["Email"] = "wkekai@163.com"
	c.TplName = "index.html"
}

func (c * MainController) ArticleInfo() {
	slug := c.Ctx.Input.Param(":slug") + ".html"

	info := models.Article{Slug: slug}
	c.o.Read(&info, "slug")
	fmt.Println(info)

	// get tag list
	var tags orm.ParamsList
	tagsId := strings.Split(info.TagIds, ",")

	qs := c.o.QueryTable(new(models.Tag))
	qs.Filter("id__in", tagsId).ValuesFlat(&tags, "name")

	// get category info
	var category models.Category
	qs = c.o.QueryTable(new(models.Category))
	qs.Filter("id", info.CategoryId).One(&category, "name")

	c.Data["info"] = info
	c.Data["tags"] = tags
	c.Data["category"] = category.Name
	c.Data["Title"] = info.Title
	c.TplName = "article.html"
}

func (this *MainController) PageInfo() {
	slug := this.Ctx.Input.Param(":slug")

	info := models.PageInfo{Slug: slug}
	this.o.Read(&info, "slug")

	this.Data["info"] = info
	this.Data["Title"] = slug

	this.TplName = "page.html"
}
