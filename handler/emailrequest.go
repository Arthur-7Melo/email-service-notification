package handler

type EmailRequest struct {
	Email   string `json:"email" binding:"required"`
	Subject string `json:"subject" binding:"required,max=100"`
	Body    string `json:"body" binding:"required,max=200"`
}