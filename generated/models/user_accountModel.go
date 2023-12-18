package models

import "time"

type UserAccount struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	Login string `json:"login"`
	PersonId string `json:"person_id"`
	Email string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Password string `json:"password"`
	Salt string `json:"salt"`
	ResetPasswordCode string `json:"reset_password_code"`
	ResetPasswordExpiration time.Time `json:"reset_password_expiration"`
	LastLogin time.Time `json:"last_login"`
	LastAccess time.Time `json:"last_access"`
	IsSuperuser bool `json:"is_superuser"`
	UserTypeId string `json:"user_type_id"`
	ReceiveSystemNotification bool `json:"receive_system_notification"`
  }