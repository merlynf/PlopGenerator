package models

import "time"

type MessageType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	MessageTypeCode string `json:"message_type_code"`
	MessageTypeName string `json:"message_type_name"`
	Priority int64 `json:"priority"`
  }