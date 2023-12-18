package models

import "time"

type EmployeeAdditionalRole struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EmployeeId string `json:"employee_id"`
	JobTitleId string `json:"job_title_id"`
	DivisionId string `json:"division_id"`
  }