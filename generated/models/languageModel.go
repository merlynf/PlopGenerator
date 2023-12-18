package models

import "time"

type Language struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	LanguageCode string `json:"language_code"`
	LanguageName string `json:"language_name"`
	LanguageNativeName string `json:"language_native_name"`
  }