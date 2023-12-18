package models

import "time"

type Customer struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	CustomerName string `json:"customer_name"`
	ProfileType int64 `json:"profile_type"`
	ProfileStatus int64 `json:"profile_status"`
	BusinessEntityId string `json:"business_entity_id"`
	IndustryId string `json:"industry_id"`
	ReferralSourceId string `json:"referral_source_id"`
	ReferralSourceName string `json:"referral_source_name"`
	Unqualified bool `json:"unqualified"`
	Description string `json:"description"`
  }