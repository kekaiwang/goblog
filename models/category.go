package models

import "time"

type Category struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	RouterLink  string    `json:"router_link"`
	LinkArticle int       `json:"link_article"`
	Status      int       `json:"status"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}
