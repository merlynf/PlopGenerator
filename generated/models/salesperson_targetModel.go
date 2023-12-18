package models

import "time"

type SalespersonTarget struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EmployeeId string `json:"employee_id"`
	HotelId string `json:"hotel_id"`
	Year string `json:"year"`
	Month string `json:"month"`
	CurrencyId string `json:"currency_id"`
	RecurringRevenue interface{} `json:"recurring_revenue"`
	LeadCalls interface{} `json:"lead_calls"`
  }