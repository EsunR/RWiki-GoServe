package controllers

import (
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"RWiki-GoServe/utils"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// @router /register [post]
func (c *BaseController) Register() {
	var resp _struct.Resp
	body := utils.GetBody(c.Ctx)
	var user models.Users
	user.Name = body["name"]
	user.Password = body["password"]
	user.Identity = "user"
	tokenString, err := models.AddOneUser(&user)
	if err != nil {
		resp.Msg = "用户名已被占用"
	} else {
		resp.Msg = "ok"
		resp.Data = map[string]interface{}{
			"token": tokenString,
		}
	}
	_ = c.Ctx.Output.JSON(resp, false, false)
}
