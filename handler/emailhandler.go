package handler

import (
	"github.com/Arthur-7Melo/email-service-notification.git/config"
	apierror "github.com/Arthur-7Melo/email-service-notification.git/config/apiError"
	"github.com/Arthur-7Melo/email-service-notification.git/config/logger"
	"github.com/Arthur-7Melo/email-service-notification.git/service"
	"github.com/gin-gonic/gin"
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

	}
}