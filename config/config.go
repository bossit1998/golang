package config

import (
	"github.com/spf13/cast"
	"os"
)

// Config ...
type Config struct {
	Environment       string // develop, staging, production
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	LogLevel          string
	RPCPort           string
	UserServiceHost   string
	UserServicePort   int
	ReviewServiceHost string
	ReviewServicePort int
	CronSetting       string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "api.test.fiesta.jafton.com"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5434))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "fiestadb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "fiesta"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "fiesta123"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9003"))
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "127.0.0.1"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 1234))
	c.ReviewServiceHost = cast.ToString(getOrReturnDefault("REVIEW_SERVICE_HOST", "127.0.0.1"))
	c.ReviewServicePort = cast.ToInt(getOrReturnDefault("REVIEW_SERVICE_PORT", 1236))
	c.CronSetting = cast.ToString(getOrReturnDefault("CRON_VALUE", "@hourly"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
