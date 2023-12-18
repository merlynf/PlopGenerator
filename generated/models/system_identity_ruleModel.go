package models

import "time"

type SystemIdentityRule struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemId string `json:"system_id"`
	IdentityCode string `json:"identity_code"`
	IdentityName string `json:"identity_name"`
	HotelGroupId string `json:"hotel_group_id"`
	HotelId string `json:"hotel_id"`
	BusinessUnitId string `json:"business_unit_id"`
	RecruitmentSystemId string `json:"recruitment_system_id"`
	Prefix string `json:"prefix"`
	DynamicPrefix string `json:"dynamic_prefix"`
	NextNumber string `json:"next_number"`
	IsReset bool `json:"is_reset"`
	ResetFrequency interface{} `json:"reset_frequency"`
	LastNumber string `json:"last_number"`
	LastUsedCounter int64 `json:"last_used_counter"`
	LastYear int64 `json:"last_year"`
	LastMonth int64 `json:"last_month"`
	Description string `json:"description"`
  }