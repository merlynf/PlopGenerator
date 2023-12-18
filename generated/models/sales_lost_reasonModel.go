package models

import "time"

type SalesLostReason struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	ReasonCode string `json:"reason_code"`
	ReasonName string `json:"reason_name"`
  }