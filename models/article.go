package models

import "time"

type Article struct {
	Id          int       `json:"id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Count       int       `json:"count"`
	Markdown    string    `json:"markdown"`
	Content     string    `json:"content"`
	CategoryId  int       `json:"category_id"`
	TagIds      string    `json:"tag_ids"`
	Excerpt     string    `json:"excerpt"`
	Previous    string    `json:"previous"`
	Next        string    `json:"next"`
	Preview     int       `json:"preview"`
	Thumb       string    `json:"thumb"`
	Slug        string    `json:"slug"`
	IsDraft     int       `json:"is_draft"`
	Created     time.Time `json:"created"`
	EditTime    time.Time `json:"edit_time"`
	Updated     time.Time `json:"updated"`
	DisplayTime time.Time `json:"display_time"`
}
