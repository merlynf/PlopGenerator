package models

import "time"

type SalesProspect struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelId string `json:"hotel_id"`
	ProspectDate time.Time `json:"prospect_date"`
	CustomerId string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	ContactName string `json:"contact_name"`
	ContactMobileNumber string `json:"contact_mobile_number"`
	ContactPhoneNumber string `json:"contact_phone_number"`
	ContactEmail string `json:"contact_email"`
	ReferralSourceId string `json:"referral_source_id"`
	ReferralSourceName string `json:"referral_source_name"`
	ProspectStatus int64 `json:"prospect_status"`
	CustomerProfileStatus int64 `json:"customer_profile_status"`
	ContactProfileStatus int64 `json:"contact_profile_status"`
	AssignById string `json:"assign_by_id"`
	AssignToTeamId string `json:"assign_to_team_id"`
	AssignToId string `json:"assign_to_id"`
	EventTypeId string `json:"event_type_id"`
	EventTypeName string `json:"event_type_name"`
	NumberOfRooms interface{} `json:"number_of_rooms"`
	NumberOfPersons interface{} `json:"number_of_persons"`
	CurrencyId string `json:"currency_id"`
	EstimatedRevenue interface{} `json:"estimated_revenue"`
	ActualRevenue interface{} `json:"actual_revenue"`
	UnqualifiedReasonId string `json:"unqualified_reason_id"`
	UnqualifiedReasonName string `json:"unqualified_reason_name"`
	LostReasonId string `json:"lost_reason_id"`
	LostReasonName string `json:"lost_reason_name"`
	Remark string `json:"remark"`
	Description string `json:"description"`
  }