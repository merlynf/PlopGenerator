package models

import "time"

type Module struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	ModuleCode string `json:"module_code"`
	ModuleName string `json:"module_name"`
	ModulePackageId string `json:"module_package_id"`
	Url string `json:"url"`
	AlwaysIncluded bool `json:"always_included"`
	Description string `json:"description"`
  }