package models

import "time"

type TaskStatus struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	StatusCode string `json:"status_code"`
	StatusName string `json:"status_name"`
	InternalStatusName string `json:"internal_status_name"`
	ConsumerStatusName string `json:"consumer_status_name"`
  }