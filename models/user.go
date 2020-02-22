package models

import (
	"net/http"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/wkekai/goblog/helper"
)

type AdminUser struct {
	Id         int       `orm:"auto"`
	Name       string    `orm:"size(255)"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	LoginCount int       `json:"login_count"`
	Salt       string    `json:"salt"`
	Status     int8      `json:"status"`
	LastLogin  time.Time `json:"last_login"`
	LastIp     string    `json:"last_ip"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

type LoginRes struct {
	Status int
	Id     int
	Token  string
}

func Login(name, password, ip string) *LoginRes {
	o := orm.NewOrm()
	user := AdminUser{Name: name}
	o.Read(&user, "name")

	if user.Password == "" {
		return &LoginRes{http.StatusNoContent, 0, ""}
	}

	if user.Password != helper.Md5(user.Name, password, user.Salt) {
		return &LoginRes{http.StatusUnauthorized, user.Id, ""}
	}

	user.LastIp = ip
	user.LastLogin = time.Now()
	user.LoginCount += 1

	o.Update(&user)

	return &LoginRes{http.StatusOK, user.Id, user.Password}
}
