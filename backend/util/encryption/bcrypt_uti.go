package encryption

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) string{
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}