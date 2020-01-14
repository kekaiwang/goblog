package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/wkekai/goblog/helper"
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
	CategoryLink string
}

type ArticleList struct {
	Data interface{}
	Total int64
}

func (c *MainController) Get() {
	var (
		articleList ArticleList
		articleInfos []ArticleInfo
		pagesize int = 5
		page int
		offset int
	)

	if page, _ = c.GetInt("page"); page < 1 {
		page = 1
	}

	offset = (page - 1) * pagesize

	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select(
		"article.*",
		"category.name as category_name",
		"category.router_link as category_link").
		From("wkk_article as article").
		LeftJoin("wkk_category as category").On("article.category_id = category.id").
		Where("article.is_draft < 3").
		OrderBy("article.id").Desc().
		Limit(pagesize).Offset(offset).String()

	c.o.Raw(sql).QueryRows(&articleInfos)

	articleList.Total, _ = c.o.QueryTable(new(models.Article)).Filter("is_draft", 2).Count()
	articleList.Data = articleInfos

	c.Data["Data"] = articleList.Data
	c.Data["Total"] = articleList.Total
	c.Data["CurrentPage"] = page
	c.Data["TotalPage"] = helper.NewTotalPage(articleList.Total, pagesize)
	c.Data["Title"] = "blog"
	c.Data["Email"] = "wkekai@163.com"
	c.TplName = "index.html"
}

func (c * MainController) ArticleInfo() {
	slug := c.Ctx.Input.Param(":slug") + ".html"

	info := models.Article{Slug: slug}
	c.o.Read(&info, "slug")

	if info.Id <= 0 {

	}

	DoArticleSum(info.Id)

	// get previous or next
	previousTitle := models.Article{Slug: info.Previous}
	if info.Previous != "" {
		c.o.QueryTable(new(models.Article)).Filter("slug", previousTitle.Slug).One(&previousTitle, "title")
	}

	nextTitle := models.Article{Slug: info.Next}
	if info.Next != "" {
		c.o.QueryTable(new(models.Article)).Filter("slug", nextTitle.Slug).One(&nextTitle, "title")
	}

	if info.Next != "" {
		c.o.QueryTable(new(models.Article)).Filter("slug", previousTitle.Slug).One(&previousTitle, "title")
	}

	// get tag list
	var tags []*models.Tag
	tagsId := strings.Split(info.TagIds, ",")

	qs := c.o.QueryTable(new(models.Tag))
	qs.Filter("id__in", tagsId).All(&tags)

	// get category info
	var category models.Category
	qs = c.o.QueryTable(new(models.Category))
	qs.Filter("id", info.CategoryId).One(&category, "name", "router_link")

	c.Data["info"] = info
	c.Data["tags"] = tags
	c.Data["category"] = category
	c.Data["previousTitle"] = previousTitle.Title
	c.Data["nextTitle"] = nextTitle.Title
	c.Data["Title"] = info.Title
	c.Data["Description"] = info.Excerpt
	c.TplName = "article.html"
}

func DoArticleSum(id int) {
	article := models.NewArticleId(id)
	models.ArticleIdM.Ch <- article
}

func (this *MainController) PageInfo() {
	slug := this.Ctx.Input.Param(":slug")

	info := models.PageInfo{Slug: slug}
	this.o.Read(&info, "slug")

	this.Data["info"] = info
	this.Data["Title"] = slug

	this.TplName = "page.html"
}

func (this *MainController) Archives() {

}

type CategoryInfo struct {
	Id int
	Name string
	Title string
	Slug string
	DisplayTime time.Time
}

type Meta struct {
	Type string
	Name string
}

func (this *MainController) Categories() {
	var (
		meta Meta
		data []CategoryInfo
		total int64
		pagesize int = 6
		page int
		offset int
		uris []string
		link string
	)

	uris = strings.Split(this.Ctx.Input.URL(), "/")

	link = this.Ctx.Input.Param(":link") + ".html"

	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}

	offset = (page - 1) * pagesize

	meta.Type = this.GetString("type")

	if uris[1] == "categories" {
		meta.Name = "分类"

		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select(
			"a.id",
			"a.name",
			"b.title",
			"b.slug",
			"b.display_time",
		).
			From("wkk_category as a").
			LeftJoin("wkk_article as b").On("a.id = b.category_id").
			Where("a.router_link = ?").
			And("b.is_draft < 3").
			OrderBy("b.id").Desc().
			Limit(pagesize).Offset(offset).String()

		this.o.Raw(sql, link).QueryRows(&data)

		this.o.Raw("select count(*) as total from wkk_category a left join wkk_article b on a.id = b.category_id " +
			"where a.router_link = ? and b.is_draft < 3", link).QueryRow(&total)
	} else if uris[1] == "tags" {
		meta.Name = "标签"

		qb, _ := orm.NewQueryBuilder("mysql")
		sql := qb.Select(
			"c.id",
			"a.name",
			"c.title",
			"c.slug",
			"c.display_time",
		).
			From("wkk_tag as a").
			LeftJoin("wkk_article_relation as b").On("a.id = b.tag_id").
			LeftJoin("wkk_article as c").On("b.article_id = c.id").
			Where("a.router_link = ?").
			And("c.is_draft < 3").
			OrderBy("c.id").Desc().
			Limit(pagesize).Offset(offset).String()

		this.o.Raw(sql, link).QueryRows(&data)

		this.o.Raw("select count(*) as total from wkk_tag a left join wkk_article_relation b on a.id = b.tag_id " +
			"left join wkk_article c on b.article_id = c.id where a.router_link = ? and c.is_draft < 3", link).QueryRow(&total)
	}

	this.Data["Data"] = data
	this.Data["Total"] = total
	this.Data["CurrentPage"] = page
	this.Data["TotalPage"] = helper.NewTotalPage(total, pagesize)
	this.Data["meta"] = meta
	this.Data["Title"] = "分类"
	this.TplName = "category.html"
}
