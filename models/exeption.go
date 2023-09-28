package models

import "time"

type Exception struct {
	Error          string    `json:"error"`
	Message        string    `json:"message"`
	Status         int       `json:"status"`
	TimeStamp      time.Time `json:"timeStamp"`
	DevelopMessage string    `json:"developMessage"`
	TransactionId  string    `json:"transactionId"`
	CorrelationId  string    `json:"correlationId"`
}
