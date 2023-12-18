package models

import "time"

type SalespersonAnalytics struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EmployeeId string `json:"employee_id"`
	MetricsId string `json:"metrics_id"`
	BusinessUnitId string `json:"business_unit_id"`
	CurrencyId string `json:"currency_id"`
	Value string `json:"value"`
	DayOfTheMonth string `json:"day_of_the_month"`
	Month string `json:"month"`
	Year string `json:"year"`
  }