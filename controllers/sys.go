package controllers

import (
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SysController struct {
	beego.Controller
}

// @router /sysStatus [get]
func (c *SysController) SysStatus() {

}

// @router /sysInit [post]
func (c *SysController) SysInit() {
	password := c.GetString("password")
	var resp _struct.Resp
	o := orm.NewOrm()
	var adminList []models.Users
	_, _ = o.QueryTable("users").Filter("identity", "admin").All(&adminList)
	if len(adminList) != 0 {
		resp.Msg = "系统已初始化，请勿重复操作"
	} else {
		var user models.Users
		user.Name = "admin"
		user.Identity = "admin"
		user.Password = password
		tokenString, err := models.AddOneUser(&user)
		if err != nil {
			fmt.Println(err)
		}
		userInfo, _ := models.GetUserInfo(user.Id)
		resp.Data = map[string]interface{}{
			"userInfo": userInfo,
			"token":    tokenString,
		}
	}
	_ = c.Ctx.Output.JSON(resp, false, false)
}
