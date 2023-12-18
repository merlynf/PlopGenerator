package models

import "time"

type HotelSalesperson struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelId string `json:"hotel_id"`
	EmployeeId string `json:"employee_id"`
	CurrencyId string `json:"currency_id"`
	RecurringRevenue interface{} `json:"recurring_revenue"`
	HasSeasonalRecurringRevenue bool `json:"has_seasonal_recurring_revenue"`
	LeadCalls interface{} `json:"lead_calls"`
	HasSeasonalLeadCalls bool `json:"has_seasonal_lead_calls"`
	TeamId string `json:"team_id"`
	RevenueSourceId string `json:"revenue_source_id"`
	AssistantId string `json:"assistant_id"`
	AssignPersonalAccount bool `json:"assign_personal_account"`
	AssignAllCorporates bool `json:"assign_all_corporates"`
	AssignAllIndustries bool `json:"assign_all_industries"`
	AssignAllCities bool `json:"assign_all_cities"`
  }