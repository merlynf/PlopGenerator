package models

import "time"

type EventType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EventTypeCode string `json:"event_type_code"`
	EventTypeName string `json:"event_type_name"`
  }