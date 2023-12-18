package models

import "time"

type System struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SystemName string `json:"system_name"`
	BusinessEntityId string `json:"business_entity_id"`
	TermsAndConditions string `json:"terms_and_conditions"`
	PrivacyPolicy string `json:"privacy_policy"`
	FromEmail string `json:"from_email"`
	FromEmailDisplay string `json:"from_email_display"`
	ToEmail string `json:"to_email"`
	ToEmailDisplay string `json:"to_email_display"`
	SessionTimeoutMinutes int64 `json:"session_timeout_minutes"`
	PasswordMinLength int64 `json:"password_min_length"`
	PasswordRequireLowercase bool `json:"password_require_lowercase"`
	PasswordRequireUppercase bool `json:"password_require_uppercase"`
	PasswordRequireDigit bool `json:"password_require_digit"`
	PasswordRequireSpecialCharacter bool `json:"password_require_special_character"`
	PasswordExpiryDays int64 `json:"password_expiry_days"`
	PasswordHistoryCheck int64 `json:"password_history_check"`
	PasswordResetLinkExpiryMinutes int64 `json:"password_reset_link_expiry_minutes"`
	MaxLoginAttempts int64 `json:"max_login_attempts"`
	LockoutDurationMinutes int64 `json:"lockout_duration_minutes"`
	BusinessUnitId string `json:"business_unit_id"`
	LoyaltyProgramId string `json:"loyalty_program_id"`
	Description string `json:"description"`
  }