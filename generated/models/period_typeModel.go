package models

import "time"

type PeriodType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	PeriodTypeCode string `json:"period_type_code"`
	PeriodTypeName string `json:"period_type_name"`
  }