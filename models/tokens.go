package models

import (
	"RWiki-GoServe/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Tokens struct {
	Id   string `orm:"pk"` // id 需要手动创建，其内容是一个时间戳
	User *Users `orm:"rel(fk)"`
}

// 传入 uid 创建一个 Token 并返回
func CreateTokenByUser(user *Users) (string, error) {
	// 检查已有的 Token 是否已经达到极限
	o := orm.NewOrm()
	var tokens []*Tokens
	qs := o.QueryTable("tokens").RelatedSel()
	_, err := qs.All(&tokens)
	if err != nil {
		return "", err
	}
	maxDevices, _ := strconv.Atoi(beego.AppConfig.String("maxDevices"))
	if len(tokens) >= maxDevices {
		oldest := tokens[0]
		_, _ = o.Delete(&Tokens{Id: oldest.Id})
	}

	// 生成 Token
	tid := strconv.FormatInt(time.Now().Unix(), 10)
	tokenString := utils.GenerateToken(map[string]interface{}{
		"uid": user.Id,
		"tid": tid,
	})

	// 插入 Token 到数据库
	tokenData := Tokens{
		Id:   tid,
		User: user,
	}
	_, err = o.Insert(&tokenData)
	if err != nil {
		return "", err
	}

	// 处理完成返回 Token
	return tokenString, nil
}
