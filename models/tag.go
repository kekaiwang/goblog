package models

import "time"

type Tag struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	RouterLink string    `json:"router_link"`
	UseTimes   int8      `json:"use_times"`
	Status     int       `json:"status"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}
