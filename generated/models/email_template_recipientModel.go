package models

import "time"

type EmailTemplateRecipient struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EmailTemplateId string `json:"email_template_id"`
	AttentionType string `json:"attention_type"`
	ToEmail string `json:"to_email"`
	ToDisplay string `json:"to_display"`
	ToPersonId string `json:"to_person_id"`
	ToEmployeeId string `json:"to_employee_id"`
	ToUserAccountId string `json:"to_user_account_id"`
  }