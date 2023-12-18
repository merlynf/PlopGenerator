package models

import "time"

type MultimediaDescription struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	InformationType string `json:"information_type"`
	PictureCategory string `json:"picture_category"`
	Title string `json:"title"`
	DimensionCategory string `json:"dimension_category"`
	FileName string `json:"file_name"`
	FileSize interface{} `json:"file_size"`
	OriginalFileName string `json:"original_file_name"`
	Url string `json:"url"`
	Description string `json:"description"`
  }