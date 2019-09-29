package models

import (
	"RWiki-GoServe/utils"
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Users struct {
	Id           int `orm:"auto"`
	Identity     string
	Name         string    `orm:"unique"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Tokens       []*Tokens `orm:"reverse(many)"`
	Password     string
	Projects     []*Projects `orm:"reverse(many)"`
	JoinProjects []*Projects `orm:"reverse(many)"`
}

// 返回所有的用户
func SelectAllUsers() ([]*Users, error) {
	o := orm.NewOrm()
	var users []*Users
	_, err := o.QueryTable("users").All(&users)
	if err != nil {
		return make([]*Users, 0), err
	}
	return users, nil
}

// 添加一个新用户 返回创建该用户后自动创建的 token
func AddOneUser(user *Users) (string, error) {
	// 对密码加密
	user.Password = utils.PwdEncode(user.Password)
	o := orm.NewOrm()
	_, err := o.Insert(user)
	if err != nil {
		return "", err
	}
	tokenString, err := CreateTokenByUser(user)
	return tokenString, err
}

// 由用户 id 获取用户的基本信息
func GetUserInfo(id int) (map[string]interface{}, error) {
	var user Users
	user.Id = id
	o := orm.NewOrm()
	err := o.Read(&user)
	if err != nil {
		return map[string]interface{}{}, errors.New("查找不到该用户的任何信息")
	}
	result := map[string]interface{}{
		"_id":      user.Id,
		"identity": user.Identity,
		"name":     user.Name,
		"time":     user.Created,
	}
	return result, nil
}
