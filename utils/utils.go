package utils

import (
	"encoding/json"
	context2 "github.com/astaxie/beego/context"
)

// 传入上下文获取请求体中的数据
func GetBody(ctx *context2.Context) map[string]interface{} {
	var body map[string]interface{}
	_ = json.Unmarshal(ctx.Input.RequestBody, &body)
	return body
}
