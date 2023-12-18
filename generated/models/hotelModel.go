package models

import "time"

type Hotel struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemId string `json:"system_id"`
	HotelCode string `json:"hotel_code"`
	HotelName string `json:"hotel_name"`
	ProfileStatus int64 `json:"profile_status"`
	EffectiveDate time.Time `json:"effective_date"`
	ExpireDate time.Time `json:"expire_date"`
	BusinessEntityId string `json:"business_entity_id"`
	HotelTypeId string `json:"hotel_type_id"`
	SalesSegmentation bool `json:"sales_segmentation"`
	SalesTeamwork bool `json:"sales_teamwork"`
	YearlyRecurringRevenueIncreasePercent interface{} `json:"yearly_recurring_revenue_increase_percent"`
	YearlyLeadCallsIncreasePercent interface{} `json:"yearly_lead_calls_increase_percent"`
	Overview string `json:"overview"`
	Directions string `json:"directions"`
	Description string `json:"description"`
  }