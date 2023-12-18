package models

import "time"

type HotelGroupTypeUserType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelGroupTypeId string `json:"hotel_group_type_id"`
	UserTypeId string `json:"user_type_id"`
	IsDefault bool `json:"is_default"`
  }