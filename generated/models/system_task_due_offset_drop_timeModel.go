package models

import "time"

type SystemTaskDueOffsetDropTime struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemId string `json:"system_id"`
	TaskDueOffsetDropTimeId string `json:"task_due_offset_drop_time_id"`
  }