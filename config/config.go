package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment      string // develop, staging, production
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	RPCPort          string
	FareServiceHost  string
	FareServicePort  int
	CronSetting      string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "127.0.0.1"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))

	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "deleverdb"))

	//c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "delever2"))
	//c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "delever"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "delever"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "delever"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9003"))
	c.FareServiceHost = cast.ToString(getOrReturnDefault("FARE_SERVICE_HOST", "127.0.0.1"))
	c.FareServicePort = cast.ToInt(getOrReturnDefault("FARE_SERVICE_PORT", 8004))
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
