package config

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	RedisAddress     string
	JWTSecretKey     string
}
