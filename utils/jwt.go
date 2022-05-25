package utils

import (
	"ByteDance/config"
	"ByteDance/pkg/msg"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"time"
)

/*
JWT使用
*/

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 这里额外记录username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

//密钥
var MySecret = []byte(config.MySecret)

/**
生成 Token
*/

func GenToken(username string) (string, error) {
	c := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.TokenExpirationTime)), // 过期时间2小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                 // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                 // 生效时间
		}}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

/**
解析 Token
*/
func ParseToken(tokenStr string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(msg.TokenValidationErrorMalformed)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(msg.TokenValidationErrorExpired)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(msg.TokenValidationErrorNotValidYet)
			} else {
				return nil, errors.New(msg.TokenHandleFailed)
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	//失效的token
	return nil, errors.New(msg.TokenValid)
}

/**
测试
*/
func main() {

	//tokenStr, _ := GenToken("徐先生")
	//fmt.Println("token:", tokenStr)
	claim, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IuW-kOWFiOeUnyIsImlzcyI6InhoeCIsImV4cCI6MTY1MzQ0OTc3MiwibmJmIjoxNjUzNDQ5NzcxLCJpYXQiOjE2NTM0NDk3NzF9.xjuR-Z39M_f_NqWRTtGjRtPBCwxS7JeaqQyDmnF7om8")
	CatchErr("错误", err)
	fmt.Printf("解析后：%#v\n", claim.ExpiresAt)
	//tokenStr2, err := RefreshToken(tokenStr)
	//
	//CatchErr("错误",err)
	//
	//fmt.Println("refToken:", tokenStr2)

}