package models

import "time"

type Division struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	DivisionName string `json:"division_name"`
	ParentDivisionId string `json:"parent_division_id"`
	ManagerId string `json:"manager_id"`
	Sales bool `json:"sales"`
	Housekeeping bool `json:"housekeeping"`
	Maintenance bool `json:"maintenance"`
	Depth int64 `json:"depth"`
  }