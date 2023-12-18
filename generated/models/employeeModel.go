package models

import "time"

type Employee struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EmployeeNumber string `json:"employee_number"`
	PersonId string `json:"person_id"`
	UserAccountId string `json:"user_account_id"`
	JobTitleId string `json:"job_title_id"`
	HireDate time.Time `json:"hire_date"`
	StartDate time.Time `json:"start_date"`
	TerminationDate time.Time `json:"termination_date"`
	DivisionId string `json:"division_id"`
	EmploymentTypeId string `json:"employment_type_id"`
	JobPositionChangeReasonId string `json:"job_position_change_reason_id"`
	JobPositionChangeReasonName string `json:"job_position_change_reason_name"`
  }