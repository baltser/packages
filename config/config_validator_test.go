package config

import (
	"strings"
	"testing"
)

func TestEnsureRequiredFields(t *testing.T) {
	testCases := []struct {
		name     string
		config   *Config
		expected string
	}{
		{
			name: "all fields filled",
			config: &Config{
				PostgresHost:     "localhost",
				PostgresPort:     5432,
				PostgresUser:     "myuser",
				PostgresPassword: "mypassword",
				PostgresDBName:   "mydb",
				RedisAddress:     "127.0.0.1:6379",
				JWTSecretKey:     "mysecretkey",
			},
			expected: "",
		},
		{
			name: "missing PostgresHost",
			config: &Config{
				PostgresPort:     5432,
				PostgresUser:     "myuser",
				PostgresPassword: "mypassword",
				PostgresDBName:   "mydb",
				RedisAddress:     "127.0.0.1:6379",
				JWTSecretKey:     "mysecretkey",
			},
			expected: "POSTGRESHOST",
		},
		{
			name: "missing multiple fields",
			config: &Config{
				PostgresHost: "localhost",
				PostgresUser: "myuser",
				RedisAddress: "127.0.0.1:6379",
				JWTSecretKey: "mysecretkey",
			},
			expected: "POSTGRESPORT, POSTGRESPASSWORD, POSTGRESDBNAME",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.ensureRequiredFields()
			if tc.expected == "" {
				if err != nil {
					t.Errorf("ensureRequiredFields() returned an unexpected error: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("ensureRequiredFields() did not return an expected error")
				} else {
					if !strings.Contains(err.Error(), tc.expected) {
						t.Errorf("ensureRequiredFields() returned an unexpected error message, expected to contain '%s', got '%s'", tc.expected, err.Error())
					}
				}
			}
		})
	}
}
