package models

import "time"

type Session struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SessionKey string `json:"session_key"`
	ExpireAt time.Time `json:"expire_at"`
	UserAccountId string `json:"user_account_id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	DeviceId string `json:"device_id"`
	Payload string `json:"payload"`
  }