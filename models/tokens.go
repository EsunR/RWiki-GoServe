package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type Tokens struct {
	Id   int64  `orm:"pk"` // id 需要手动创建，其内容是一个时间戳
	User *Users `orm:"rel(fk)"`
}

// 传入 uid 创建一个 Token 并返回
func CreateToken(user *Users) (string, error) {
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
	claims := make(jwt.MapClaims)
	tid := time.Now().Unix()
	claims["uid"] = user.Id
	claims["tid"] = tid
	jwtExpiresTime, _ := strconv.Atoi(beego.AppConfig.String("jwtExpiresTime"))
	claims["exp"] = time.Now().Add(time.Duration(jwtExpiresTime) * time.Millisecond).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := beego.AppConfig.String("jwtSecret")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", nil
	}

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
