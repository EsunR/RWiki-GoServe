package filters

import (
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"RWiki-GoServe/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"net/http"
	"regexp"
	"strings"
)

// 全局变量存放 Token 信息
var TokenData map[string]interface{}
var BodyData map[string]interface{}

func init() {
	patter := "register|login|sysInit|test"
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
		// CORS
		if context.Request.Method == http.MethodOptions {
			context.Output.Status = http.StatusOK
			context.WriteString("ok")
		}

		// 身份鉴权
		_ = json.Unmarshal(context.Input.RequestBody, &BodyData)
		if match, _ := regexp.MatchString(patter, context.Request.RequestURI); !match {
			resp := _struct.Resp{}
			defer func() {
				// 错误处理（如果未携带Token）
				r := recover()
				if err, ok := r.(error); ok {
					fmt.Println(err)
					resp.Msg = "请求未携带 Token，请尝试重新登录"
					context.Output.Status = http.StatusForbidden
					_ = context.Output.JSON(resp, false, false)
				}
			}()
			auth := context.Input.Header("Authorization")
			tokenString := strings.Split(auth, " ")[1]
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
				} else {
					TokenData = tokenData
				}
			}

		}
	})
}
