package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	Db     DbConfig
	Auth   AuthConfig
}

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	return &Config{
		Server: ServerConfig{
			Port: os.Getenv("APPLICATION_PORT"),
		},
		Db: DbConfig{
			Dsn: dsn,
		},
		Auth: AuthConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
