package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// CryptPasswd encrypt passwd string to hash string
func CryptPasswd(passwd string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	return string(hashed)
}

// ComparePasswd compare with hashed string
func ComparePasswd(passwd, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwd))
	if err != nil {
		return false
	}
	return true
}
