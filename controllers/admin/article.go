package admin

import (
	"encoding/json"
	"github.com/wkekai/goblog/RS"
	"github.com/wkekai/goblog/helper"
	"github.com/wkekai/goblog/models"
	"time"
)

type ArticleController struct {
	baseController
}

// -------------- tag --------------------
type TagInfo struct {
	Data []*models.Tag
	Total int64
}

func (article *ArticleController) TagList() {
	resp := helper.NewResponse()
	var info TagInfo

	limit, _ := article.GetInt("limit")
	page, _ := article.GetInt("page", 0)
	page = (page - 1) * limit

	query := article.o.QueryTable(new(models.Tag))

	query.OrderBy("-id").Limit(limit, page).All(&info.Data)

	info.Total, _ = query.Count()
	resp.Data = info

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) CreateTag() {
	var tag models.Tag
	resp := helper.NewResponse()

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	json.Unmarshal(article.Ctx.Input.RequestBody, &tag)

	query := article.o.QueryTable(new(models.Tag))
	classifyNum, _ := query.Filter("name", tag.Name).Count()

	if classifyNum > 0 {
		resp.Status = RS.RS_tag_exist
		resp.Tips(helper.WARNING, "名称已存在")
		return
	}

	tag.Created = time.Now()
	_, err := article.o.Insert(&tag)

	if err != nil {
		resp.Status = RS.RS_create_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_tag_exist))
		return
	}

	resp.Data = tag
}

func (article *ArticleController) UpdateTag() {
	resp := helper.NewResponse()
	var tag models.Tag

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	json.Unmarshal(article.Ctx.Input.RequestBody, &tag)

	info := models.Tag{Id: tag.Id}

	article.o.Read(&info, "id")

	if info.Name == "" {
		resp.Status = RS.RS_not_found
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_not_found))
		return
	}

	info.Name = tag.Name
	info.Status = tag.Status
	info.Updated = time.Now()

	if num, err := article.o.Update(&info); err != nil {
		resp.Status = RS.RS_update_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_update_failed))
		return
	} else {
		resp.Data = num
	}
}

// -------------- category --------------------
type CategoryList struct {
	Data []*models.Category
	Total int64
}

func (article *ArticleController) CategoryList() {
	resp := helper.NewResponse()

	var categoryList CategoryList

	limit, _ := article.GetInt("limit")
	page, _ := article.GetInt("page")
	start := (page - 1) * limit

	query := article.o.QueryTable(new(models.Category))
	query.OrderBy("-id").Limit(limit, start).All(&categoryList.Data)

	categoryList.Total, _ = query.Count()
	resp.Data = categoryList

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) CreateCategory() {
	resp := helper.NewResponse()
	var category models.Category

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	json.Unmarshal(article.Ctx.Input.RequestBody, &category)

	query := article.o.QueryTable(new(models.Category))
	categoryNum, _ := query.Filter("name", category.Name).Count()

	if categoryNum > 0 {
		resp.Status = RS.RS_tag_exist
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_tag_exist))
		return
	}

	category.Created = time.Now()

	_, err := article.o.Insert(&category)

	if err != nil {
		resp.Status = RS.RS_create_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_create_failed))
		return
	}

	resp.Data = category
}

func (article *ArticleController) UpdateCategory() {
	resp := helper.NewResponse()

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	var category models.Category

	json.Unmarshal(article.Ctx.Input.RequestBody, &category)

	info := models.Category{Id: category.Id}
	article.o.Read(&info, "id")

	if info.Name == "" {
		resp.Status = RS.RS_not_found
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_not_found))
		return
	}

	info.Name = category.Name
	info.Status = category.Status
	info.Updated = time.Now()

	if num, err := article.o.Update(&info); err != nil {
		resp.Status = RS.RS_update_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_update_failed))
		return
	} else {
		resp.Data = num
	}
}

// -------------- article --------------------
type ArticleList struct {
	Data []*models.Article
	Total int64
}

func (article *ArticleController) ArticleList() {
	resp := helper.NewResponse()
	var articleList ArticleList

	limit, _ := article.GetInt("limit")
	page, _ := article.GetInt("page")
	start := (page - 1) * limit

	query := article.o.QueryTable(new(models.Article))
	query.OrderBy("-id").Limit(limit, start).All(&articleList.Data)

	articleList.Total, _ = query.Count()
	resp.Data = articleList

	resp.WriteJson(article.Ctx.ResponseWriter)
}
