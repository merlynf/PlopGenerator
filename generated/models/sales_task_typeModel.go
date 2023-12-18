package models

import "time"

type SalesTaskType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SalesTaskTypeCode string `json:"sales_task_type_code"`
	SalesTaskTypeName string `json:"sales_task_type_name"`
  }