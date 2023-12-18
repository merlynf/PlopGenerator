package models

import "time"

type UserThrottle struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	UserAccountId string `json:"user_account_id"`
	IpAddress string `json:"ip_address"`
	LoginAttempts interface{} `json:"login_attempts"`
	LastAttemptAt time.Time `json:"last_attempt_at"`
	Requests interface{} `json:"requests"`
	RequestUntil time.Time `json:"request_until"`
	IsSuspended bool `json:"is_suspended"`
	SuspendedAt time.Time `json:"suspended_at"`
	IsBanned bool `json:"is_banned"`
	BannedAt time.Time `json:"banned_at"`
	Comment string `json:"comment"`
  }