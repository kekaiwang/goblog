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

type TagInfo struct {
	Data []*models.Tag
	Total int64
}

func (article *ArticleController) ClassifyList() {
	resp := helper.NewResponse()
	var info TagInfo

	//page := article.GetString("page")
	//limit := article.GetString("limit")

	query := article.o.QueryTable(new(models.Tag))

	query.OrderBy("-id").Limit(10, 0).All(&info.Data)

	info.Total, _ = query.Count()
	resp.Data = info

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) CreateClassify() {
	var classify models.Tag
	resp := helper.NewResponse()

	json.Unmarshal(article.Ctx.Input.RequestBody, &classify)

	query := article.o.QueryTable(new(models.Tag))
	classifyNum, _ := query.Filter("name", classify.Name).Count()

	if classifyNum > 0 {
		resp.Status = RS.RS_tag_exist
		resp.Tips(helper.WARNING, "名称已存在")
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	}

	classify.Created = time.Now()
	_, err := article.o.Insert(&classify)

	if err != nil {
		resp.Status = RS.RS_create_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_tag_exist))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	}

	resp.Data = classify

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) UpdateClassify() {
	resp := helper.NewResponse()
	var classify models.Tag

	json.Unmarshal(article.Ctx.Input.RequestBody, &classify)

	tag := models.Tag{Id: classify.Id}

	article.o.Read(&tag, "id")

	if tag.Name == "" {
		resp.Status = RS.RS_not_found
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_not_found))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	}

	tag.Name = classify.Name
	tag.Status = classify.Status
	tag.Updated = time.Now()

	if num, err := article.o.Update(&tag); err != nil {
		resp.Status = RS.RS_update_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_update_failed))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	} else {
		resp.Data = num
	}

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) ModifyClassify() {
	resp := helper.NewResponse()

	var classify models.Tag

	json.Unmarshal(article.Ctx.Input.RequestBody, &classify)

	classify.Updated = time.Now()

	if num, err := article.o.Update(&classify, "status", "updated"); err != nil {
		resp.Status = RS.RS_update_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_update_failed))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	} else {
		resp.Data = num
	}

	resp.WriteJson(article.Ctx.ResponseWriter)
}
