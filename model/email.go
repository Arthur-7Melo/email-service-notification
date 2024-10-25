package model

import "time"

type EmailNotification struct {
	Id      string `bson:"_id,omitempty"`
	Email   string `bson:"email"`
	Subject string `bson:"subject"`
	Body    string `bson:"body"`
	Sent_At time.Time `bson:"sent_at"`
}