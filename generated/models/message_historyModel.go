package models

import "time"

type MessageHistory struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	MessageId string `json:"message_id"`
	MessageStatus int64 `json:"message_status"`
	UserAccountId string `json:"user_account_id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	DeviceId string `json:"device_id"`
	OperatingSystemId string `json:"operating_system_id"`
	Comment string `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
	LocalTimestamp time.Time `json:"local_timestamp"`
  }