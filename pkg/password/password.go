package password

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword хэширует пароль с использованием соли
func HashPassword(password string) (string, string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", "", err
	}
	saltStr := base64.StdEncoding.EncodeToString(salt)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password+saltStr), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return string(hashedBytes), saltStr, nil
}

// CheckPassword проверяет соответствие пароля, соли и хэша
func CheckPassword(password, salt, hash string) bool {
	passwordWithSalt := password + salt

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordWithSalt))
	return err == nil
}
