package models

import "time"

type ProfileStatus struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	ProfileStatusCode string `json:"profile_status_code"`
	ProfileStatusName string `json:"profile_status_name"`
  }