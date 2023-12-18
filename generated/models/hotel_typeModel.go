package models

import "time"

type HotelType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemId string `json:"system_id"`
	HotelTypeCode string `json:"hotel_type_code"`
	HotelTypeName string `json:"hotel_type_name"`
	IsDefault bool `json:"is_default"`
  }