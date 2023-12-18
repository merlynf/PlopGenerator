package models

import "time"

type HotelGroupJobTitle struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelGroupId string `json:"hotel_group_id"`
	JobTitleId string `json:"job_title_id"`
	PositionOpenDate time.Time `json:"position_open_date"`
	TargetStartDate time.Time `json:"target_start_date"`
	JobStatus int64 `json:"job_status"`
  }