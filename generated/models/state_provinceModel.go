package models

import "time"

type StateProvince struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	StateProvinceCode string `json:"state_province_code"`
	StateProvinceName string `json:"state_province_name"`
	CountryId string `json:"country_id"`
  }