package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) string {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while encrypting the password ", err)
		return ""
	}
	return string(hash)
}
