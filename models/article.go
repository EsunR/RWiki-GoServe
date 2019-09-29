package models

import "time"

type Articles struct {
	Id         int `orm:"auto"`
	Title      string
	Md         string
	Html       string
	Project    *Projects `orm:"rel(fk)"`
	CreateTime time.Time `orm:"auto_now_add;type(date)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}
