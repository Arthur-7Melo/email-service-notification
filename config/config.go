package config

import (
	"log"
	"os"

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

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil{
		log.Fatal("Erro ao carregar variáveis de ambiente .env")
	}

	return &Config{
		MongoURI: os.Getenv("MONGODB_URI"),
		MongoDB: os.Getenv("MONGODB_DB"),
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: os.Getenv("SMTP_PORT"),
		SMTPUser: os.Getenv("SMTP_USER"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
}