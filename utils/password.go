package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	
	if err != nil {
		return "", err
	}

	return string(byte), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}