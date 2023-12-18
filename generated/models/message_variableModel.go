package models

import "time"

type MessageVariable struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	MessageVariableCode string `json:"message_variable_code"`
	MessageVariableName string `json:"message_variable_name"`
	DummyValue string `json:"dummy_value"`
	Description string `json:"description"`
  }