package models

import "time"

type HotelRevenueSource struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	RevenueSourceCode string `json:"revenue_source_code"`
	RevenueSourceName string `json:"revenue_source_name"`
	IsDefault bool `json:"is_default"`
  }