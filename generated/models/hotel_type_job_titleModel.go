package models

import "time"

type HotelTypeJobTitle struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelTypeId string `json:"hotel_type_id"`
	JobTitleId string `json:"job_title_id"`
  }