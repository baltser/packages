package config

import (
	"github.com/lpernett/godotenv"
	"os"
	"strconv"
	"sync"
)

var once sync.Once

func (c *Config) Getting() error {

	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		c.PostgresHost = os.Getenv(envPostgresHost)
		c.PostgresPort, _ = strconv.Atoi(os.Getenv(envPostgresPort))
		c.PostgresUser = os.Getenv(envPostgresUser)
		c.PostgresPassword = os.Getenv(envPostgresPassword)
		c.PostgresDBName = os.Getenv(envPostgresDBName)
		c.RedisAddress = os.Getenv(envRedisAddress)
		c.JWTSecretKey = os.Getenv(envJWTSecretKey)
	})

	return c.ensureRequiredFields()
}
