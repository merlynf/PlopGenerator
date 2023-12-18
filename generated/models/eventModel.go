package models

import "time"

type Event struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EventCode string `json:"event_code"`
	EventName string `json:"event_name"`
	EventTypeId string `json:"event_type_id"`
	EventLevelId string `json:"event_level_id"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
  }