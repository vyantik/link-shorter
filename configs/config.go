package configs

import (
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

	return &Config{
		Server: ServerConfig{
			Port: os.Getenv("APPLICATION_PORT"),
		},
		Db: DbConfig{
			Dsn: os.Getenv("POSTGRES_URI"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SESSION_SECRET"),
		},
	}
}
