package models

import "time"

type PageInfo struct {
	Id int
	Name string
	Slug string
	Content string
	Markdown string
	Preview int
	Status int
	Created time.Time
	Updated time.Time
}
