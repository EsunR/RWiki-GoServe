package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// 由 Token 字符串 返回一个 Token map
func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(beego.AppConfig.String("jwtSecret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func GenerateToken(payload map[string]interface{}) string {
	claims := make(jwt.MapClaims)
	for k, v := range payload {
		claims[k] = v
	}
	jwtExpiresTime, _ := strconv.Atoi(beego.AppConfig.String("jwtExpiresTime"))
	claims["exp"] = time.Now().Add(time.Duration(jwtExpiresTime) * time.Millisecond).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := beego.AppConfig.String("jwtSecret")
	tokenString, _ := token.SignedString([]byte(secret))
	return "Bearer " + tokenString
}
