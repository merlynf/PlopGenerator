package models

import "time"

type NamePrefix struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	NamePrefixCode string `json:"name_prefix_code"`
	NamePrefixName string `json:"name_prefix_name"`
	GenderId string `json:"gender_id"`
  }