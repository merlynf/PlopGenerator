package models

import "time"

type Mail struct {
	Id string `json:"id"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	RecepientMail string `json:"recepient_mail"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Remark string `json:"remark"`
  }