package models

import "time"

type CustomerAssignment struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	HotelId string `json:"hotel_id"`
	CustomerId string `json:"customer_id"`
	IndustryId string `json:"industry_id"`
	CityId string `json:"city_id"`
	AssignToTeamId string `json:"assign_to_team_id"`
	AssignToId string `json:"assign_to_id"`
  }