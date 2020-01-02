package models

import "time"

type AdminUser struct {
	Id int `orm:"auto"`
	Name string `orm:"size(255)"`
	Password string
	Email string
	LoginCount int
	status int8
	LastLogin time.Time
	LastIp string
	Created time.Time
	Updated time.Time
}
