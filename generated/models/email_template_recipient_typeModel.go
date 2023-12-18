package models

import "time"

type EmailTemplateRecipientType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EmailTemplateId string `json:"email_template_id"`
	RecipientTypeId string `json:"recipient_type_id"`
	AttentionType string `json:"attention_type"`
  }