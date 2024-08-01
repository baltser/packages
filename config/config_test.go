package config

import (
	"os"
	"strings"
	"testing"
)

func TestGetting(t *testing.T) {
	// Arrange
	os.Setenv(envPostgresHost, "localhost")
	os.Setenv(envPostgresPort, "5432")
	os.Setenv(envPostgresUser, "myuser")
	os.Setenv(envPostgresPassword, "mypassword")
	os.Setenv(envPostgresDBName, "mydb")
	os.Setenv(envRedisAddress, "127.0.0.1:6379")
	os.Setenv(envJWTSecretKey, "mysecretkey")
	defer func() {
		os.Unsetenv(envPostgresHost)
		os.Unsetenv(envPostgresPort)
		os.Unsetenv(envPostgresUser)
		os.Unsetenv(envPostgresPassword)
		os.Unsetenv(envPostgresDBName)
		os.Unsetenv(envRedisAddress)
		os.Unsetenv(envJWTSecretKey)
	}()

	config := &Config{}

	// Act
	err := config.Getting()

	// Assert
	if err != nil {
		t.Errorf("Getting() returned an error: %v", err)
	}

	if config.PostgresHost != "localhost" {
		t.Errorf("Expected PostgresHost to be 'localhost', got '%s'", config.PostgresHost)
	}

	if config.PostgresPort != 5432 {
		t.Errorf("Expected PostgresPort to be 5432, got %d", config.PostgresPort)
	}

	if config.PostgresUser != "myuser" {
		t.Errorf("Expected PostgresUser to be 'myuser', got '%s'", config.PostgresUser)
	}

	if config.PostgresPassword != "mypassword" {
		t.Errorf("Expected PostgresPassword to be 'mypassword', got '%s'", config.PostgresPassword)
	}

	if config.PostgresDBName != "mydb" {
		t.Errorf("Expected PostgresDBName to be 'mydb', got '%s'", config.PostgresDBName)
	}

	if config.RedisAddress != "127.0.0.1:6379" {
		t.Errorf("Expected RedisAddress to be '127.0.0.1:6379', got '%s'", config.RedisAddress)
	}

	if config.JWTSecretKey != "mysecretkey" {
		t.Errorf("Expected JWTSecretKey to be 'mysecretkey', got '%s'", config.JWTSecretKey)
	}
}

func TestGetting_MissingEnvVars(t *testing.T) {
	// Arrange
	os.Unsetenv(envPostgresHost)
	os.Unsetenv(envPostgresPort)
	config := &Config{}

	// Act
	err := config.Getting()

	// Assert
	if err == nil {
		t.Error("Getting() did not return an error")
	}

	expectedErrorMsg := "missing required environment variables"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedErrorMsg, err.Error())
	}
}
