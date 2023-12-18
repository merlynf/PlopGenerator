package models

import "time"

type PictureCategory struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	PictureCategoryCode string `json:"picture_category_code"`
	PictureCategoryName string `json:"picture_category_name"`
  }