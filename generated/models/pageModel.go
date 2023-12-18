package models

import "time"

type Page struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	PageName string `json:"page_name"`
	Slug string `json:"slug"`
	HtmlTitle string `json:"html_title"`
	PageTypeId string `json:"page_type_id"`
	PageFileName string `json:"page_file_name"`
	PublishedAt time.Time `json:"published_at"`
  }