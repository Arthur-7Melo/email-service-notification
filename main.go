package main

import (
	"context"
	"os"

	"github.com/Arthur-7Melo/email-service-notification.git/config"
	"github.com/Arthur-7Melo/email-service-notification.git/config/logger"
	"github.com/Arthur-7Melo/email-service-notification.git/db"
	"github.com/Arthur-7Melo/email-service-notification.git/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Iniciando api email-notification")
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Erro ao carregar variáveis de ambiente", err)
	}

	db.ConnectMongoDb(cfg)
	defer db.MongoClient.Disconnect(context.TODO())

	router := gin.Default()
	router.POST("/send-email", handler.SendEmailHandler(cfg))

	port := os.Getenv("PORT")
	router.Run(":" + port)
}