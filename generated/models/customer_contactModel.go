package models

import "time"

type CustomerContact struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	CustomerId string `json:"customer_id"`
	PersonId string `json:"person_id"`
	ProfileStatus int64 `json:"profile_status"`
	DivisionName string `json:"division_name"`
	IsPrimary bool `json:"is_primary"`
	Unqualified bool `json:"unqualified"`
  }