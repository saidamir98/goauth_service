package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App string

	Environment string // development, staging, production

	LogLevel    string // debug, info, warn, error, dpanic, panic, fatal
	ServiceHost string
	HTTPPort    string
	BasePath    string

	DefaultOffset string
	DefaultLimit  string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	CassandraHost     string
	CassandraPort     int
	CassandraKeyspace string
	CassandraUser     string
	CassandraPassword string

	RabbitURI string

	SecretKey string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("PROJECT_NAME", "goauth_service"))

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))
	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7071"))
	config.BasePath = cast.ToString(getOrReturnDefaultValue("BASE_PATH", "/v1"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "goauth_service"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", ""))

	config.CassandraHost = cast.ToString(getOrReturnDefaultValue("CASSANDRA_HOST", "localhost"))
	config.CassandraPort = cast.ToInt(getOrReturnDefaultValue("CASSANDRA_PORT", 9042))
	config.CassandraKeyspace = cast.ToString(getOrReturnDefaultValue("CASSANDRA_KEYSPACE", "goauth_service"))
	config.CassandraUser = cast.ToString(getOrReturnDefaultValue("CASSANDRA_USER", "cassandra"))
	config.CassandraPassword = cast.ToString(getOrReturnDefaultValue("CASSANDRA_PASSWORD", "cassandra"))

	config.RabbitURI = cast.ToString(getOrReturnDefaultValue("AMQP_URI", "amqp://guest:guest@localhost:5672"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "Here$houldBe$ome$ecretKey"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
