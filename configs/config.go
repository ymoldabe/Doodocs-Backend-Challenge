package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	Host                string
	EmailSenderName     string
	EmailSenderAddres   string
	EmailSenderPassword string
}

func New() (Config, error) {
	godotenv.Load(".env")

	config := Config{
		Port:                os.Getenv("PORT"),
		Host:                os.Getenv("HOST"),
		EmailSenderName:     os.Getenv("EMAIL_SENDER_NAME"),
		EmailSenderAddres:   os.Getenv("EMAIL_SENDER_ADDRES"),
		EmailSenderPassword: os.Getenv("EMAIL_SENDER_PASSWORD"),
	}

	return config, nil
}
