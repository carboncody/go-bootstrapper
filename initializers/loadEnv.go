package initializers

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig() (config Config, err error) {
    err = godotenv.Load(".env")
    if err != nil {
        return
    }

    config = Config{
        DBHost:         os.Getenv("POSTGRES_HOST"),
        DBUserName:     os.Getenv("POSTGRES_USER"),
        DBUserPassword: os.Getenv("POSTGRES_PASSWORD"),
        DBName:         os.Getenv("POSTGRES_DB"),
        DBPort:         os.Getenv("POSTGRES_PORT"),
        ServerPort:     os.Getenv("PORT"),
        ClientOrigin:   os.Getenv("CLIENT_ORIGIN"),
        // Add other environment variables as needed
    }

    return config, nil
}

