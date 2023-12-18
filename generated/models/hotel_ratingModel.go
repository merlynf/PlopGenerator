package models

import "time"

type HotelRating struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelId string `json:"hotel_id"`
	RatingTypeId string `json:"rating_type_id"`
	Rating interface{} `json:"rating"`
  }