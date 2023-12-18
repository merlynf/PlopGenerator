package models

import "time"

type LanguageSkillLevel struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	LevelCode string `json:"level_code"`
	LevelName string `json:"level_name"`
	Level interface{} `json:"level"`
  }