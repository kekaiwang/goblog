package models

import "time"

type Article struct {
	Id int
	Author string
	Title string
	Count int
	Markdown string
	Content string
	CategoryId int
	TagIds string
	Excerpt string
	Previous string
	Next string
	Preview int
	Slug string
	IsDraft int
	Created time.Time
	EditTime time.Time
	Updated time.Time
	DisplayTime time.Time
}
