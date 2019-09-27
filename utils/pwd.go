package utils

import "golang.org/x/crypto/bcrypt"

func PwdEncode(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func PwdCompare(password string, encodePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
