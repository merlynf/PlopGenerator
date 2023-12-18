package models

import "time"

type TimeOfTheDay struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	TimeOfTheDayCode string `json:"time_of_the_day_code"`
	TimeOfTheDayName string `json:"time_of_the_day_name"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
  }