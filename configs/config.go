package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
}

func New() (Config, error) {
	err := godotenv.Load("./configs/config.env")
	if err != nil {
		return Config{}, err
	}

	config := Config{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}

	return config, nil
}
