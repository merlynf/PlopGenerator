package models

import "time"

type ModulePackage struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	ModulePackageCode string `json:"module_package_code"`
	ModulePackageName string `json:"module_package_name"`
	Description string `json:"description"`
  }