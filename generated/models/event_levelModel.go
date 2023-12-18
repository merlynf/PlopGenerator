package models

import "time"

type EventLevel struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EventLevelCode string `json:"event_level_code"`
	EventLevelName string `json:"event_level_name"`
	Level int64 `json:"level"`
  }