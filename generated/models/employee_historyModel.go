package models

import "time"

type EmployeeHistory struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EmployeeId string `json:"employee_id"`
	JobTitleId string `json:"job_title_id"`
	StartDate time.Time `json:"start_date"`
	DivisionId string `json:"division_id"`
	EmploymentTypeId string `json:"employment_type_id"`
	JobPositionChangeReasonId string `json:"job_position_change_reason_id"`
	JobPositionChangeReasonName string `json:"job_position_change_reason_name"`
  }