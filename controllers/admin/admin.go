package admin

import (
	"encoding/json"
	"github.com/wkekai/goblog/RS"
	"github.com/wkekai/goblog/helper"
	"github.com/wkekai/goblog/models"
)

type UserController struct {
	baseController
}

type UserParams struct {
	Username string
	Password string
}

type Info struct {
	Roles string
	Name string
	Avatar string
	Introduction string
}

func (admin *UserController) Login() {
	resp := helper.NewResponse()
	var user UserParams

	err := json.Unmarshal(admin.Ctx.Input.RequestBody, &user)

	if err != nil {
		resp.Tips(helper.WARNING, "123")
		resp.WriteJson(admin.Ctx.ResponseWriter)
		return
	}

	if user.Username == "" || user.Password == "" {
		resp.Status = RS.RS_params_error
		resp.Tips(helper.WARNING, "123")
		resp.WriteJson(admin.Ctx.ResponseWriter)
		return
	}

	status := models.Login(user.Username, user.Password, admin.GetClientIp())

	resp.Status = status.Status

	switch status.Status {
	case 204:
		resp.Tips(helper.ERROR, "管理员账号错误")
		resp.WriteJson(admin.Ctx.ResponseWriter)
		return
	case 401:
		resp.Tips(helper.WARNING, "管理员账号/密码错误")
		resp.WriteJson(admin.Ctx.ResponseWriter)
		return
	}

	resp.Data = status.Token

	resp.WriteJson(admin.Ctx.ResponseWriter)
}

func (admin *UserController) GetInfo() {
	resp := helper.NewResponse()

	var info Info

	token := admin.GetString("token")

	if token == "editor-token" {
		info.Roles = "editor"
	} else {
		info.Roles = "admin"
	}

	info.Name = "wkekai"
	info.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	info.Introduction = "wkekai blog"

	resp.Data = info

	resp.WriteJson(admin.Ctx.ResponseWriter)
}

func (admin *UserController) Logout() {
	resp := helper.NewResponse()

	resp.WriteJson(admin.Ctx.ResponseWriter)
}
