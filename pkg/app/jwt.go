package app

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyCustomClaims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": util.EncodeMD5(appKey),
		"id":   util.EncodeMD5(appSecret),
		"exp":  time.Now().Add(global.JWTSetting.Expire).Unix(), // 设置令牌过期时间为24小时
	})
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*MyCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		fmt.Println("ccc", err)
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
