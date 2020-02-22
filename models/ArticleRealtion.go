package models

type ArticleRelation struct {
	Id        int   `json:"id"`
	ArticleId int64 `json:"article_id"`
	TagId     int   `json:"tag_id"`
}
