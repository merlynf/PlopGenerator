package models

import "time"

type PageType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	PageTypeCode int64 `json:"page_type_code"`
	PageTypeName string `json:"page_type_name"`
  }