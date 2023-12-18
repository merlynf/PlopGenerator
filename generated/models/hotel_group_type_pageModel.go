package models

import "time"

type HotelGroupTypePage struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelGroupTypeId string `json:"hotel_group_type_id"`
	PageId string `json:"page_id"`
  }