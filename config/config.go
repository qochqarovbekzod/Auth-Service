package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT string
	GRPC_PORT string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_PASSWORD string
	DB_USER     string

	AccessToken string
	RefreshToken string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("NO .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8081"))
	config.GRPC_PORT = cast.ToString(coalesce("GRPC_PORT", ":50050"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("db_port", "5432"))
	config.DB_NAME = cast.ToString(coalesce("db_name", "users"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "1918"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.AccessToken = cast.ToString(coalesce("AccessToken", "123456"))
	config.RefreshToken = cast.ToString(coalesce("RefreshToken", "19181"))

	return config
}

func coalesce(key string, value string) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
