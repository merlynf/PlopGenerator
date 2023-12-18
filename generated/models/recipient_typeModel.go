package models

import "time"

type RecipientType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	RecipientTypeCode int64 `json:"recipient_type_code"`
	RecipientTypeName string `json:"recipient_type_name"`
	BuiltInTypeId string `json:"built_in_type_id"`
  }