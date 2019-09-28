package routers

import (
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"RWiki-GoServe/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"net/http"
	"regexp"
)

func init() {
	patter := "register|login"
	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
		if match, _ := regexp.MatchString(patter, context.Request.RequestURI); !match {
			resp := _struct.Resp{}
			tokenString := context.Input.Header("Authorization")
			if tokenString == "" {
				resp.Msg = "请求未携带Token，请尝试重新登录"
				context.Output.Status = http.StatusForbidden
				_ = context.Output.JSON(resp, false, false)
			} else {
				// 解析 Token 是否过期
				tokenData, err := utils.ParseToken(tokenString)
				if err != nil {
					resp.Msg = "登录 Token 已过期，请重新登录"
					context.Output.Status = http.StatusForbidden
					_ = context.Output.JSON(resp, false, false)
				} else {
					o := orm.NewOrm()
					var tokenModel models.Tokens
					tokenModel.Id = tokenData["tid"].(string)
					err := o.Read(&tokenModel)
					if err != nil {
						resp.Msg = "登录 Token 已失效，请重新登录"
						context.Output.Status = http.StatusForbidden
						_ = context.Output.JSON(resp, false, false)
					}
				}
			}
		}
	})
}
