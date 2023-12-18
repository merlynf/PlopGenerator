package models

import "time"

type EmailTemplate struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	MessageTypeId string `json:"message_type_id"`
	FromEmail string `json:"from_email"`
	FromDisplay string `json:"from_display"`
	FromPersonId string `json:"from_person_id"`
	FromEmployeeId string `json:"from_employee_id"`
	FromUserAccountId string `json:"from_user_account_id"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	BodyFileName string `json:"body_file_name"`
	ScheduleId string `json:"schedule_id"`
	Description string `json:"description"`
  }