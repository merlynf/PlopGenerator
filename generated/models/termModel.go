package models

import "time"

type Term struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	TermCode string `json:"term_code"`
	TermName string `json:"term_name"`
	Slug string `json:"slug"`
	ParentTermId string `json:"parent_term_id"`
  }