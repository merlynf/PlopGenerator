package models

import "time"

type SalesProspectAsset struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	SalesProspectId string `json:"sales_prospect_id"`
	MultimediaDescriptionId string `json:"multimedia_description_id"`
  }