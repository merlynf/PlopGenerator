package models

import "time"

type MessageTypeEmailRecipient struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	MessageTypeId string `json:"message_type_id"`
	AttentionType string `json:"attention_type"`
	ToEmail string `json:"to_email"`
	ToDisplay string `json:"to_display"`
  }