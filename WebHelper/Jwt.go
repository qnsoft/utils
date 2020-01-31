package WebHelper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(SecretKey []byte, issuer string, user_id int, is_admin bool) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		user_id,
		is_admin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

/**
 * 解析 校验token是否有效
 */
func CheckToken(token string, SecretKey string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		//fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}

/**
 * 解析 token
 */
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}
