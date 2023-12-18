package models

import "time"

type HotelGroupType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemId string `json:"system_id"`
	HotelGroupTypeCode string `json:"hotel_group_type_code"`
	HotelGroupTypeName string `json:"hotel_group_type_name"`
	IsDefault bool `json:"is_default"`
  }