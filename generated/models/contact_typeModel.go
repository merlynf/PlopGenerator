package models

import "time"

type ContactType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	ContactTypeCode string `json:"contact_type_code"`
	ContactTypeName string `json:"contact_type_name"`
  }