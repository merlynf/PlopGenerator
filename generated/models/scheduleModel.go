package models

import "time"

type Schedule struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ScheduleName string `json:"schedule_name"`
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	AbsoluteDeadline time.Time `json:"absolute_deadline"`
	OffsetTimeUnitId string `json:"offset_time_unit_id"`
	OffsetUnitMultiplier int64 `json:"offset_unit_multiplier"`
	OffsetDropTimeId string `json:"offset_drop_time_id"`
	RecurrenceTypeId string `json:"recurrence_type_id"`
	Seconds string `json:"seconds"`
	Minutes string `json:"minutes"`
	Hours string `json:"hours"`
	DayOfTheMonth string `json:"day_of_the_month"`
	Months string `json:"months"`
	DayOfTheWeek string `json:"day_of_the_week"`
	Years string `json:"years"`
	Description string `json:"description"`
  }