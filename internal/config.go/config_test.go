package config

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func setupEnv() {
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "5432")
	os.Setenv("DATABASE_USER", "testuser")
	os.Setenv("DATABASE_PASSWORD", "testpass")
	os.Setenv("DATABASE_NAME", "testdb")
	os.Setenv("JWT_SECRET_KEY", "secretkey")
	os.Setenv("SERVER_PORT", "9999")
}

func TestNewConfig(t *testing.T) {
	setupEnv()
	cfg := NewConfig()

	if cfg.Database.Host != "localhost" {
		t.Errorf("Ожидался Host 'localhost', получен '%s'", cfg.Database.Host)
	}
	if cfg.Server.Port != "9999" {
		t.Errorf("Ожидался порт '9999', получен '%s'", cfg.Server.Port)
	}
	if cfg.JWT.SecretKey != "secretkey" {
		t.Errorf("Ожидался JWT ключ 'secretkey', получен '%s'", cfg.JWT.SecretKey)
	}
}

func TestGetConnStr(t *testing.T) {
	db := DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		Name:     "testdb",
	}
	connStr := db.GetConnStr()

	expected := "postgresql://testuser:testpass@localhost:5432/testdb?sslmode=disable"
	if connStr != expected {
		t.Errorf("Ожидался connStr '%s', получен '%s'", expected, connStr)
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("EXISTING_ENV", "value")
	val := getEnv("EXISTING_ENV", "default")
	if val != "value" {
		t.Errorf("Ожидалось 'value', получено '%s'", val)
	}

	val = getEnv("NON_EXISTING_ENV", "default")
	if val != "default" {
		t.Errorf("Ожидалось 'default', получено '%s'", val)
	}
}

func TestGetEnvOrFatal_MissingVariable(t *testing.T) {
	if os.Getenv("BE_CRASH_TEST") == "1" {
		_ = NewConfig()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestGetEnvOrFatal_MissingVariable")
	cmd.Env = append(os.Environ(), "BE_CRASH_TEST=1")
	cmd.Env = removeEnvVars(cmd.Env, []string{
		"DATABASE_HOST", "DATABASE_PORT", "DATABASE_USER",
		"DATABASE_PASSWORD", "DATABASE_NAME", "JWT_SECRET_KEY",
	})
	output, err := cmd.CombinedOutput()

	if err == nil {
		t.Fatalf("Ожидался выход с ошибкой из-за отсутствия переменной окружения")
	}

	if !strings.Contains(string(output), "Environment variable") {
		t.Errorf("Ожидалось сообщение об ошибке, получено: %s", output)
	}
}

func removeEnvVars(env []string, keys []string) []string {
	var result []string
	for _, e := range env {
		keep := true
		for _, k := range keys {
			if strings.HasPrefix(e, k+"=") {
				keep = false
				break
			}
		}
		if keep {
			result = append(result, e)
		}
	}
	return result
}
