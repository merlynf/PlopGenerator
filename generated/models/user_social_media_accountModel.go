package models

import "time"

type UserSocialMediaAccount struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	UserAccountId string `json:"user_account_id"`
	SocialMediaId string `json:"social_media_id"`
	UserIdentifier string `json:"user_identifier"`
	Username string `json:"username"`
	Email string `json:"email"`
	OauthToken string `json:"oauth_token"`
  }