package models

import "time"

type Projects struct {
	Id           int `orm:"auto"`
	Desc         string
	Cover        string
	ProjectName  string
	Creator      *Users      `orm:"rel(fk)"`
	Contributors []*Users    `orm:"rel(m2m)"`
	Articles     []*Articles `orm:"reverse(many)"`
	CreateTime   time.Time   `orm:"auto_now_add;type(date)"`
	UpdateTime   time.Time   `orm:"auto_now;type(datetime)"`
}
