package models

import "time"

type HotelGroup struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemId string `json:"system_id"`
	HotelGroupCode string `json:"hotel_group_code"`
	HotelGroupName string `json:"hotel_group_name"`
	ProfileStatus int64 `json:"profile_status"`
	EffectiveDate time.Time `json:"effective_date"`
	ExpireDate time.Time `json:"expire_date"`
	BusinessEntityId string `json:"business_entity_id"`
	HotelGroupTypeId string `json:"hotel_group_type_id"`
	BillingType string `json:"billing_type"`
	Description string `json:"description"`
  }