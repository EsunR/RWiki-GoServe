package controllers

import (
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"RWiki-GoServe/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseController struct {
	beego.Controller
}

// @router /updateToken [get]
func (c *BaseController) UpdateToken() {
	//var resp _struct.Resp

}

// @router /register [post]
func (c *BaseController) Register() {
	var resp _struct.Resp
	body := utils.GetBody(c.Ctx)
	var user models.Users
	user.Name = body["name"].(string)
	user.Password = body["password"].(string)
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

// @router /login [post]
func (c *BaseController) Login() {
	var resp _struct.Resp

	body := utils.GetBody(c.Ctx)
	name := body["name"].(string)
	password := body["password"].(string)

	var user models.Users
	user.Name = name
	o := orm.NewOrm()
	err := o.Read(&user, "name")
	if err != nil {
		// 找不到用户名
		resp.Msg = "未找到该用户"
	} else {
		// 如果找到用户名，进行密码匹配
		if utils.PwdCompare(password, user.Password) {
			// 密码匹配成功
			resp.Msg = "ok"
			tokenString, _ := models.CreateToken(&user)
			resp.Data = map[string]string{
				"token": tokenString,
			}
		} else {
			// 密码匹配失败
			resp.Msg = "密码错误，请重试"
		}
	}
	_ = c.Ctx.Output.JSON(resp, false, false)
}
