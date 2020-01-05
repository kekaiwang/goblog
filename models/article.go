package models

import "time"

type Article struct {
	Id int
	Author string
	Title string
	Count int
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
}
