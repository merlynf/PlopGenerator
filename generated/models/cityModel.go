package models

import "time"

type City struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	CityCode string `json:"city_code"`
	CityName string `json:"city_name"`
	CountryId string `json:"country_id"`
	StateProvinceId string `json:"state_province_id"`
	TimezoneId string `json:"timezone_id"`
  }