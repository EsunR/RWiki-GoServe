package controllers

import (
	"RWiki-GoServe/filters"
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"RWiki-GoServe/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type BaseController struct {
	beego.Controller
}

// @router /updateToken [get]
func (c *BaseController) UpdateToken() {
	// 此时 Token 是一个未过期并且有效的 Token，需要检测其是否即将过期，然后更新 Token
	var resp _struct.Resp
	now := time.Now().Unix()
	tid, _ := strconv.ParseInt(filters.TokenData["tid"].(string), 10, 64)
	// TODO: Token 在这里的更新机制存在问题
	jwtExpiresTime, _ := strconv.ParseInt(beego.AppConfig.String("jwtExpiresTime"), 10, 64)
	lastTime := jwtExpiresTime - (now-tid)*1000 // 剩余时间（毫秒）
	if lastTime > 24*60*3600 {
		// 如果Token剩余时间大于24小时，则无需更新Token
		resp.Msg = "ok"
		resp.Data = map[string]interface{}{
			"update": false,
		}
	} else {
		// 如果 Token 即将过期，则更新 Token
		// 制定新的 tid
		newTid := strconv.FormatInt(now, 10)

		// 更新数据库中对应的 tid
		var tokenModel models.Tokens
		tokenModel.Id = filters.TokenData["tid"].(string)
		o := orm.NewOrm()
		_, _ = o.QueryTable("tokens").
			Filter("id", filters.TokenData["tid"].(string)).
			Update(orm.Params{"id": newTid})
		// 生成新的 Token 返回
		resp.Msg = "ok"
		resp.Data = map[string]interface{}{
			"token": utils.GenerateToken(map[string]interface{}{
				"tid": newTid,
				"uid": filters.TokenData["uid"],
			}),
		}
	}
	_ = c.Ctx.Output.JSON(resp, false, false)
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
			tokenString, _ := models.CreateTokenByUser(&user)
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

// @router /getUserInfo [get]
func (c *BaseController) GetUserInfo() {
	var resp _struct.Resp
	uid := int(filters.TokenData["uid"].(float64))
	userInfo, err := models.GetUserInfo(uid)
	if err != nil {
		resp.Msg = "未找到该用户的任何信息"
	} else {
		resp.Msg = "ok"
		resp.Data = userInfo
	}
	_ = c.Ctx.Output.JSON(resp, false, false)
}
