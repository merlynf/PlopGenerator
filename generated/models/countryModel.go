package models

import "time"

type Country struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	CountryAccessCode string `json:"country_access_code"`
	TimezoneId string `json:"timezone_id"`
	CurrencyId string `json:"currency_id"`
	LanguageId string `json:"language_id"`
	Nationality string `json:"nationality"`
	RegionId string `json:"region_id"`
  }