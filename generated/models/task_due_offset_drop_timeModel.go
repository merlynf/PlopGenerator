package models

import "time"

type TaskDueOffsetDropTime struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	OffsetDropTimeCode string `json:"offset_drop_time_code"`
	OffsetDropTimeName string `json:"offset_drop_time_name"`
	Offset string `json:"offset"`
  }