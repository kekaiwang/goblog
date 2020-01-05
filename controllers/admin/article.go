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

func (article *ArticleController) TagList() {
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

func (article *ArticleController) CreateTag() {
	var tag models.Tag
	resp := helper.NewResponse()

	json.Unmarshal(article.Ctx.Input.RequestBody, &tag)

	query := article.o.QueryTable(new(models.Tag))
	classifyNum, _ := query.Filter("name", tag.Name).Count()

	if classifyNum > 0 {
		resp.Status = RS.RS_tag_exist
		resp.Tips(helper.WARNING, "名称已存在")
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	}

	tag.Created = time.Now()
	_, err := article.o.Insert(&tag)

	if err != nil {
		resp.Status = RS.RS_create_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_tag_exist))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	}

	resp.Data = tag

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) UpdateTag() {
	resp := helper.NewResponse()
	var tag models.Tag

	json.Unmarshal(article.Ctx.Input.RequestBody, &tag)

	info := models.Tag{Id: tag.Id}

	article.o.Read(&info, "id")

	if info.Name == "" {
		resp.Status = RS.RS_not_found
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_not_found))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	}

	info.Name = tag.Name
	info.Status = tag.Status
	info.Updated = time.Now()

	if num, err := article.o.Update(&info); err != nil {
		resp.Status = RS.RS_update_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_update_failed))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	} else {
		resp.Data = num
	}

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) ModifyTag() {
	resp := helper.NewResponse()

	var tag models.Tag

	json.Unmarshal(article.Ctx.Input.RequestBody, &tag)

	tag.Updated = time.Now()

	if num, err := article.o.Update(&tag, "status", "updated"); err != nil {
		resp.Status = RS.RS_update_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_update_failed))
		resp.WriteJson(article.Ctx.ResponseWriter)
		return
	} else {
		resp.Data = num
	}

	resp.WriteJson(article.Ctx.ResponseWriter)
}

type ArticleList struct {

}

func (article *ArticleController) ArticleList() {

}
