package service

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/Arthur-7Melo/email-service-notification.git/config"
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