package models

import "time"

type Person struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	GenderId string `json:"gender_id"`
	NamePrefixId string `json:"name_prefix_id"`
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string `json:"last_name"`
	FullName string `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
	JobTitleName string `json:"job_title_name"`
	CitizenCountryId string `json:"citizen_country_id"`
	BusinessEntityId string `json:"business_entity_id"`
	Description string `json:"description"`
  }