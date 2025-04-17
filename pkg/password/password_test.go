package password

import (
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	passwords := []string{
		"password123",
		"сложныйПароль!@#$%^&*()",
		"1234567890",
		"",
	}

	for _, password := range passwords {
		hash, salt, err := HashPassword(password)
		if err != nil {
			t.Errorf("HashPassword(%q) вернул ошибку: %v", password, err)
			continue
		}

		if hash == "" {
			t.Errorf("HashPassword(%q) вернул пустой хеш", password)
		}
		if salt == "" {
			t.Errorf("HashPassword(%q) вернул пустую соль", password)
		}

		if !strings.HasPrefix(hash, "$2a$") && !strings.HasPrefix(hash, "$2b$") {
			t.Errorf("HashPassword(%q) вернул хеш без префикса bcrypt: %s", password, hash)
		}

		if !isBase64(salt) {
			t.Errorf("HashPassword(%q) вернул соль с недопустимыми символами: %s", password, salt)
		}
	}
}

func TestCheckPassword(t *testing.T) {
	password := "testPassword123"
	hash, salt, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword(%q) вернул ошибку: %v", password, err)
	}

	if !CheckPassword(password, salt, hash) {
		t.Errorf("CheckPassword(%q, %q, %q) вернул false для правильного пароля",
			password, salt, hash)
	}

	wrongPasswords := []string{
		"testPassword124",
		"TestPassword123",
		"testpassword123",
		"testPassword12",
		"testPassword1234",
		"",
	}

	for _, wrongPassword := range wrongPasswords {
		if CheckPassword(wrongPassword, salt, hash) {
			t.Errorf("CheckPassword(%q, %q, %q) вернул true для неправильного пароля",
				wrongPassword, salt, hash)
		}
	}

	wrongSalt := "aW52YWxpZHNhbHQ="
	if CheckPassword(password, wrongSalt, hash) {
		t.Errorf("CheckPassword(%q, %q, %q) вернул true для неправильной соли",
			password, wrongSalt, hash)
	}

	wrongHash := "$2a$10$" + strings.Repeat("a", 53)
	if CheckPassword(password, salt, wrongHash) {
		t.Errorf("CheckPassword(%q, %q, %q) вернул true для неправильного хеша",
			password, salt, wrongHash)
	}
}

func isBase64(s string) bool {
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '+' || c == '/' || c == '=' {
			continue
		}
		return false
	}
	return true
}
