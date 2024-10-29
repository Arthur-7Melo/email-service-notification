package service

import (
	"context"
	"fmt"
	"net/smtp"
	"strconv"
	"time"

	"github.com/Arthur-7Melo/email-service-notification.git/config"
	"github.com/Arthur-7Melo/email-service-notification.git/db"
	"github.com/Arthur-7Melo/email-service-notification.git/model"
)

func SendEmail(cfg *config.Config, to, subject, body string) error {
	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPassword, cfg.SMTPHost)
	port, _ := strconv.Atoi(cfg.SMTPPort)
	msg := []byte("To: " + to + "\r\n" +
								"Subject: " + subject + "\r\n" +
								"\r\n" +
								body + "\r\n")
	return smtp.SendMail(fmt.Sprintf("%s:%d", cfg.SMTPHost, port), auth, cfg.SMTPUser, []string{to}, msg)
}

func SaveEmail(cfg *config.Config, email model.EmailNotification) error {
	collection := db.GetEmailCollection(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, email)
	return err
}