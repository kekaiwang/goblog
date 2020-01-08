package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/wkekai/goblog/RS"
	"github.com/wkekai/goblog/helper"
	"github.com/wkekai/goblog/models"
	"os"
	"path"
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
	Data interface{}
	Total int64
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

func (article *ArticleController) ArticleList() {
	resp := helper.NewResponse()
	var articleList ArticleList
	var articleInfos []ArticleInfo

	limit, _ := article.GetInt("limit")
	page, _ := article.GetInt("page")
	start := (page - 1) * limit

	//query := article.o.QueryTable(new(models.Article))
	//
	//query.Filter("is_draft", 2).Limit(limit, start).RelatedSel().All(&articleList.Data)
	//resp.Data = articleList
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select(
		"article.*",
		"category.name as category_name").
		From("wkk_article as article").
		LeftJoin("wkk_category as category").On("article.category_id = category.id").
		Where("article.is_draft < 3").
		OrderBy("article.id").Desc().
		Limit(limit).Offset(start).String()

	// 执行 SQL 语句
	article.o.Raw(sql).QueryRows(&articleInfos)

	articleList.Total, _ = article.o.QueryTable(new(models.Article)).Filter("is_draft", 2).Count()
	articleList.Data = articleInfos
	resp.Data = articleList

	resp.WriteJson(article.Ctx.ResponseWriter)
}

func (article *ArticleController) GetCategories() {
	resp := helper.NewResponse()

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	var category []*models.Category

	query := article.o.QueryTable(new(models.Category))
	query.Filter("Status", 1).OrderBy("id").All(&category)

	resp.Data = category
}

func (article *ArticleController) GetTags() {
	resp := helper.NewResponse()

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	var tag []*models.Tag

	query := article.o.QueryTable(new(models.Tag))
	query.Filter("Status", 1).OrderBy("id").All(&tag)

	resp.Data = tag
}

func (article *ArticleController) Detail() {
	resp := helper.NewResponse()

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	var articleInfo models.Article
	articleInfo.Id, _ = article.GetInt("id")
	article.o.Read(&articleInfo)

	resp.Data = articleInfo
}

func (article *ArticleController) CreateArticle() {
	resp := helper.NewResponse()
	defer resp.WriteJson(article.Ctx.ResponseWriter)

	var articleInfo models.Article
	json.Unmarshal(article.Ctx.Input.RequestBody, &articleInfo)
	fmt.Println(articleInfo)

	articleInfo.Created = time.Now()

	id, err := article.o.Insert(&articleInfo)
	fmt.Println(err)
	if err != nil {
		resp.Status = RS.RS_create_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_create_failed))
		return
	}

	resp.Data = id
}

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func (article *ArticleController) UploadImage() {
	resp := helper.NewResponse()

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	file, information, err := article.GetFile("image")

	if err != nil {
		resp.Status = RS.RS_params_error
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_params_error))
		return
	}
	defer file.Close()

	localFile := information.Filename

	article.SaveToFile("image", path.Join("static/temp_image/", localFile))

	bucket := "blog-wkk"
	key := helper.Md5(localFile, helper.RandString(10), "")
	accessKey := beego.AppConfig.String("AK")
	secretKey := beego.AppConfig.String("SK")

	fmt.Println(localFile, bucket, key)
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, "static/temp_image/" + localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Remove("static/temp_image/" + localFile)

	resp.Data = ret.Key
}

func (article *ArticleController) UpdateArticle() {
	resp := helper.NewResponse()

	defer resp.WriteJson(article.Ctx.ResponseWriter)

	var articleInfo models.Article
	json.Unmarshal(article.Ctx.Input.RequestBody, &articleInfo)

	info := models.Article{Id: articleInfo.Id}
	article.o.Read(&info, "id")

	if info.Title == "" {
		resp.Status = RS.RS_not_found
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_not_found))
		return
	}

	info.Title = articleInfo.Title
	info.CategoryId = articleInfo.CategoryId
	info.DisplayTime = articleInfo.DisplayTime
	info.TagIds = articleInfo.TagIds
	info.Slug = articleInfo.Slug
	info.Content = articleInfo.Content
	info.Markdown = articleInfo.Markdown
	info.IsDraft = articleInfo.IsDraft
	info.Updated = time.Now()

	if num, err := article.o.Update(&info); err != nil {
		resp.Status = RS.RS_update_failed
		resp.Tips(helper.WARNING, RS.Desc(RS.RS_update_failed))
		return
	} else {
		resp.Data = num
	}
}
