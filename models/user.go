package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/wkekai/goblog/helper"
	"net/http"
	"time"
)

type AdminUser struct {
	Id int `orm:"auto"`
	Name string `orm:"size(255)"`
	Password string
	Email string
	LoginCount int
	Salt string
	status int8
	LastLogin time.Time
	LastIp string
	Created time.Time
	Updated time.Time
}

type LoginRes struct {
	Status int
	Id int
	Token string
}

func Login(name, password, ip string) *LoginRes {
	o := orm.NewOrm()
	user := AdminUser{Name: name}
	o.Read(&user, "name")
	fmt.Println(user)

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

