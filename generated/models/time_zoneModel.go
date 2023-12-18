package models

import "time"

type TimeZone struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	UtcOffset string `json:"utc_offset"`
	UtcOffsetHours interface{} `json:"utc_offset_hours"`
	UtcDaylightSavingTimeOffset string `json:"utc_daylight_saving_time_offset"`
	UtcDaylightSavingTimeOffsetHours interface{} `json:"utc_daylight_saving_time_offset_hours"`
	ZoneName string `json:"zone_name"`
	Description string `json:"description"`
  }