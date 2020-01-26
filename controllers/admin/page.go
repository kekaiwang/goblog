package admin

import (
	"encoding/json"
	"github.com/wkekai/goblog/helper"
	"github.com/wkekai/goblog/models"
	"time"
)

type PageController struct {
	baseController
}

type PageInfo struct {
	Data []*models.PageInfo
	Total int64
}

func (this *PageController) PageList() {
	resp := helper.NewResponse()
	defer resp.WriteJson(this.Ctx.ResponseWriter)

	var info PageInfo

	limit, _ := this.GetInt("limit")
	page, _ := this.GetInt("page")
	start := (page - 1) * limit

	query := this.o.QueryTable(new(models.PageInfo))

	query.Filter("status", 1).OrderBy("-id").Limit(limit, start).All(&info.Data)

	info.Total, _ = query.Filter("id", 1).Count()

	resp.Data = info
}

func (page *PageController) CreatePage() {
	resp := helper.NewResponse()

	defer resp.WriteJson(page.Ctx.ResponseWriter)

	var info models.PageInfo
	json.Unmarshal(page.Ctx.Input.RequestBody, &info)

	pageNum, _ := page.o.QueryTable(new(models.PageInfo)).Filter("name", info.Name).Count()

	if pageNum > 0 {
		resp.Status = helper.RS_tag_exist
		resp.Tips(helper.WARNING, helper.Desc(helper.RS_tag_exist))
		return
	}

	info.Created = time.Now()
	info.Status = 1

	_, err := page.o.Insert(&info)

	if err != nil {
		resp.Status = helper.RS_create_failed
		resp.Tips(helper.WARNING, helper.Desc(helper.RS_create_failed))
		return
	}

	resp.Data = info
}

func (page *PageController) Detail() {
	resp := helper.NewResponse()
	defer resp.WriteJson(page.Ctx.ResponseWriter)

	var info models.PageInfo
	info.Slug = page.GetString("slug")
	page.o.Read(&info, "slug")

	resp.Data = info
}

func (page *PageController) UpdatePage() {
	resp := helper.NewResponse()
	defer resp.WriteJson(page.Ctx.ResponseWriter)

	var pageInfo models.PageInfo
	json.Unmarshal(page.Ctx.Input.RequestBody, &pageInfo)

	info := models.PageInfo{Id: pageInfo.Id}
	page.o.Read(&info)

	if info.Name == "" {
		resp.Status = helper.RS_not_found
		resp.Tips(helper.WARNING, helper.Desc(helper.RS_not_found))
		return
	}

	pageInfo.Updated = time.Now()

	if num, err := page.o.Update(&pageInfo); err != nil {
		resp.Status = helper.RS_update_failed
		resp.Tips(helper.WARNING, helper.Desc(helper.RS_update_failed))
		return
	} else {
		resp.Data = num
	}
}
