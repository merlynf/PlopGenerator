package models

import "time"

type SalesEventType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SalesEventTypeCode string `json:"sales_event_type_code"`
	SalesEventTypeName string `json:"sales_event_type_name"`
  }