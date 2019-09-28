package models

import (
	"RWiki-GoServe/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GlobalController struct {
	beego.Controller
}


func init() {
	utils.DbLink()
	orm.RegisterModel(new(Tokens), new(Users))
	_ = orm.RunSyncdb("default", false, false)
}
