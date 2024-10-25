package db

import (
	"context"
	"time"

	"github.com/Arthur-7Melo/email-service-notification.git/config"
	"github.com/Arthur-7Melo/email-service-notification.git/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDb(cfg *config.Config){
	logger.Info("Iniciando conexão com o MongoDB")
	ctx, cancelDB := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelDB()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		logger.Error("Erro ao conectar com o MongoDB", err)
		return
	}

	MongoClient = client
	logger.Info("Conectado ao MongoDB")
}

func GetEmailCollection(cfg *config.Config) *mongo.Collection{
	return MongoClient.Database(cfg.MongoDB).Collection("emails")
}
