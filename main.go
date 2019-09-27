package main

import (
	_ "RWiki-GoServe/models"
	_ "RWiki-GoServe/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
