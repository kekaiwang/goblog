package models

import "time"

type Tag struct {
	Id         int
	Name       string
	RouterLink string `orm:"unique"`
	UseTimes   int8
	Status     int
	Created    time.Time
	Updated    time.Time
}
