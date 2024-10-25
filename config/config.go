package config

import (
	"os"

	"github.com/Arthur-7Melo/email-service-notification.git/config/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI     string
	MongoDB      string
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
}

func LoadConfig() (*Config, error) {
	logger.Info("Iniciando carregamento de variáveis de ambiente")
	if err := godotenv.Load(); err != nil{
		logger.Error("Erro ao carregar variáveis de ambiente .env", err)
		return nil, err
	}

	return &Config{
		MongoURI: os.Getenv("MONGODB_URI"),
		MongoDB: os.Getenv("MONGODB_DB"),
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: os.Getenv("SMTP_PORT"),
		SMTPUser: os.Getenv("SMTP_USER"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}, nil
}