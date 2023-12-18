package models

import "time"

type RatingType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	RatingTypeCode string `json:"rating_type_code"`
	RatingTypeName string `json:"rating_type_name"`
	Provider string `json:"provider"`
	Scale int64 `json:"scale"`
	RatingSymbol string `json:"rating_symbol"`
  }