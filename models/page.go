package models

import "time"

type PageInfo struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	Content  string    `json:"content"`
	Markdown string    `json:"markdown"`
	Preview  int       `json:"previcew"`
	Status   int       `json:"status"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
