package handler

import (
	"net/http"
	"time"

	"github.com/Arthur-7Melo/email-service-notification.git/config"
	apierror "github.com/Arthur-7Melo/email-service-notification.git/config/apiError"
	"github.com/Arthur-7Melo/email-service-notification.git/config/logger"
	"github.com/Arthur-7Melo/email-service-notification.git/model"
	"github.com/Arthur-7Melo/email-service-notification.git/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SendEmailHandler(cfg *config.Config) gin.HandlerFunc{

	return func(ctx *gin.Context) {
		var requestBody EmailRequest
		if err := ctx.ShouldBindJSON(&requestBody); err != nil{
			logger.Error("Erro na requisição Json", err)
			emailErr := apierror.NewBadRequestError("Erro na requisição Json")
			ctx.JSON(emailErr.Code, emailErr)
			return
		}

		if err := service.SendEmail(cfg, requestBody.Email, requestBody.Subject, requestBody.Body); err != nil{
			logger.Error("Falha ao enviar email", err)
			emailErr := apierror.NewInternalServerError("Falha ao enviar email")
			ctx.JSON(emailErr.Code, emailErr)
			return
		}

		email := model.EmailNotification{
			Email: requestBody.Email,
			Subject: requestBody.Subject,
			Body: requestBody.Body,
			Sent_At: time.Now(),
		}

		if err := service.SaveEmail(cfg, email); err != nil {
			logger.Error("Erro ao salvar email no db", err)
			emailErr := apierror.NewInternalServerError("Erro ao salvar registro de email no banco de dados")
			ctx.JSON(emailErr.Code, emailErr)
			return
		}

		logger.Info("Email enviado e salvo no banco de dados", zap.String(
			"email", requestBody.Email,
		))
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Email enviado com sucesso",
		})
	}
}