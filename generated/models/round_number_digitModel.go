package models

import "time"

type RoundNumberDigit struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	RoundDigitCode string `json:"round_digit_code"`
	RoundDigitName string `json:"round_digit_name"`
  }