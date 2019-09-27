package utils

import (
	"github.com/astaxie/beego/orm"
)

func DbLink() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/rwiki?charset=utf8")
}
