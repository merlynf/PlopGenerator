package models

import "time"

type MessageTypeVariable struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	MessageTypeId string `json:"message_type_id"`
	MessageVariableId string `json:"message_variable_id"`
  }