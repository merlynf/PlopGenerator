package models

import "time"

type UserType struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	UserTypeName string `json:"user_type_name"`
	IsSystem bool `json:"is_system"`
	AssignAllModules bool `json:"assign_all_modules"`
  }