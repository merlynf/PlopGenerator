package models

import "time"

type EventLog struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	EventId string `json:"event_id"`
	SessionId string `json:"session_id"`
	ModuleId string `json:"module_id"`
	PageId string `json:"page_id"`
	UserAccountId string `json:"user_account_id"`
	MemberId string `json:"member_id"`
	GuestCriteriaId string `json:"guest_criteria_id"`
	ReservationId string `json:"reservation_id"`
	MessageId string `json:"message_id"`
	CurrencyConversionId string `json:"currency_conversion_id"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
	Payload string `json:"payload"`
  }