package php_self

import (
	"golang.org/x/crypto/bcrypt"
)

/*
php特有的加密算法
*/
func Password_hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

/*
php特有的加密算法验证
*/
func Password_verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
