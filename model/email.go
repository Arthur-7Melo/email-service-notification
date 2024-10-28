package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailNotification struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email   string `json:"email" bson:"email"`
	Subject string `json:"subject" bson:"subject"`
	Body    string `json:"body" bson:"body"`
	Sent_At time.Time `json:"sent_at" bson:"sent_at"`
}