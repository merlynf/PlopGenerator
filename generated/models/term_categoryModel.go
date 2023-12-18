package models

import "time"

type TermCategory struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	TermCategoryCode string `json:"term_category_code"`
	TermCategoryName string `json:"term_category_name"`
	ParentTermCategoryId string `json:"parent_term_category_id"`
  }