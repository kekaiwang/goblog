package admin

import (
	"encoding/json"
	"fmt"
	"github.com/wkekai/goblog/RS"
	"github.com/wkekai/goblog/helper"
	"github.com/wkekai/goblog/models"
	"net/http"
)

type UserController struct {
	baseController
}

type UserAdmin struct {
	Username string
	Password string
}

func (admin *UserController) Login() {
	resp := helper.NewResponse()
	var user UserAdmin

	err := json.Unmarshal(admin.Ctx.Input.RequestBody, &user)

	if err != nil {
		resp.Tips(helper.WARNING, RS.RS_params_error)
		resp.WriteJson(admin.Ctx.ResponseWriter)
		return
	}

	if user.Username == "" || user.Password == "" {
		resp.Status = RS.RS_params_error
		resp.Tips(helper.WARNING, RS.RS_params_error)
		resp.WriteJson(admin.Ctx.ResponseWriter)
		return
	}

	u := models.AdminUser{Name: user.Username}
	admin.o.Read(&u, "name")

	if u.Password != "" {
		resp.Status = http.StatusNoContent
		resp.Tips(helper.ALERT, RS.RS_not_found)
		resp.WriteJson(admin.Ctx.ResponseWriter)
		return
	}

	fmt.Printf("Err: %v\n", err)

	resp.WriteJson(admin.Ctx.ResponseWriter)
}
