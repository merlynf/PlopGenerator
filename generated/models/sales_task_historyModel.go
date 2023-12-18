package models

import "time"

type SalesTaskHistory struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SalesTaskId string `json:"sales_task_id"`
	HotelId string `json:"hotel_id"`
	TaskTypeId string `json:"task_type_id"`
	TaskDate time.Time `json:"task_date"`
	CustomerId string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	ContactName string `json:"contact_name"`
	ContactMobileNumber string `json:"contact_mobile_number"`
	ContactPhoneNumber string `json:"contact_phone_number"`
	ContactEmail string `json:"contact_email"`
	ReferralSourceId string `json:"referral_source_id"`
	ReferralSourceName string `json:"referral_source_name"`
	DueDate time.Time `json:"due_date"`
	TaskStatus int64 `json:"task_status"`
	AssignById string `json:"assign_by_id"`
	AssignToTeamId string `json:"assign_to_team_id"`
	AssignToId string `json:"assign_to_id"`
	OriginatingTaskId string `json:"originating_task_id"`
	PreviousTaskId string `json:"previous_task_id"`
	SalesProspectId string `json:"sales_prospect_id"`
	ReferralTaskId string `json:"referral_task_id"`
	Result string `json:"result"`
	Remark string `json:"remark"`
	Description string `json:"description"`
  }