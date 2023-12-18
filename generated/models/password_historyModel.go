package models

import "time"

type PasswordHistory struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	UserAccountId string `json:"user_account_id"`
	IpAddress string `json:"ip_address"`
	Password string `json:"password"`
	Salt string `json:"salt"`
	LastChange time.Time `json:"last_change"`
	Expiration time.Time `json:"expiration"`
  }