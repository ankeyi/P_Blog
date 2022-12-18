package main

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	user_info
	jwt.RegisteredClaims
}

var MySecret []byte

// 获取Key,Main函数中完成了初始化判断
func init() {
	key, _ := os.ReadFile("key")
	MySecret = key
}

// 生成 JWT
func MakeToken(u_c user_info) (tokenString string, err error) {
	claim := MyClaims{
		user_info: u_c,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "kaige", //签发人
		},
	}
	// 使用指定签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定secret 签名并获得完整编码后的字符串token
	tokenString, err = token.SignedString(MySecret)
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	}
}

func ParseToken(tokenss string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}

	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
