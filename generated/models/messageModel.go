package models

import "time"

type Message struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	MessageTypeId string `json:"message_type_id"`
	MessageStatus int64 `json:"message_status"`
	Priority int64 `json:"priority"`
	FromEmail string `json:"from_email"`
	FromDisplay string `json:"from_display"`
	FromPersonId string `json:"from_person_id"`
	FromEmployeeId string `json:"from_employee_id"`
	FromUserAccountId string `json:"from_user_account_id"`
	ReplyToEmail string `json:"reply_to_email"`
	ReplyToDisplay string `json:"reply_to_display"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	BodyFileName string `json:"body_file_name"`
  }