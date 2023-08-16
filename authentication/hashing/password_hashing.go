package authentication

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil
}

func IsHashedPasswordMatch(hashedPassword string, originalPassword string) bool {
	byteHash := []byte(hashedPassword)
	byteOriginal := []byte(originalPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, byteOriginal)
	if err != nil {
		fmt.Println("hashing matcher error : ", err)
		return false
	}

	return true
}
