package models

import "time"

type DateOfTheMonth struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	DateOfTheMonthCode string `json:"date_of_the_month_code"`
	DateOfTheMonthName string `json:"date_of_the_month_name"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
  }