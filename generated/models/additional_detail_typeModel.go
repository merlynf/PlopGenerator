package models

import "time"

type AdditionalDetailType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	AdditionalDetailTypeCode string `json:"additional_detail_type_code"`
	AdditionalDetailTypeName string `json:"additional_detail_type_name"`
	Description string `json:"description"`
  }