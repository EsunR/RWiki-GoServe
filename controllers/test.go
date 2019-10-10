package controllers

import (
	"RWiki-GoServe/models"
	_struct "RWiki-GoServe/struct"
	"fmt"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) URLMapping() {
	c.Mapping("TestRouter", c.TestRouter)
}

// @router /testRouter [get]
func (c *TestController) TestRouter() {
	data := map[string]interface{}{
		"name": "hua hua",
		"age":  18,
		"arr":  []interface{}{"233", 18},
	}
	var resp _struct.Resp
	resp.Msg = "ok"
	resp.Data = data
	_ = c.Ctx.Output.JSON(resp, true, false)
}

// @router /testUsers [get]
func (c *TestController) TestUsers() {
	resp := _struct.Resp{}
	result, err := models.SelectAllUsers()
	if err != nil {
		resp.Msg = "数据库查询失败"
	} else {
		resp.Msg = "ok"
		resp.Data = result
	}
	_ = c.Ctx.Output.JSON(resp, false, false)
}

// @router /testData [get]
func (c *TestController) TestData() {
	fmt.Println("test")
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user := User{
		Name: "huahua",
		Age:  18,
	}
	c.Data["json"] = user
	c.ServeJSON()
}
