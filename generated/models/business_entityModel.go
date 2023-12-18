package models

import "time"

type BusinessEntity struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	LanguageId string `json:"language_id"`
	CurrencyId string `json:"currency_id"`
	TimeZoneId string `json:"time_zone_id"`
	AddressLine string `json:"address_line"`
	CityId string `json:"city_id"`
	CityName string `json:"city_name"`
	CountryId string `json:"country_id"`
	PostalCode string `json:"postal_code"`
	StateProvinceId string `json:"state_province_id"`
	UnformattedAddress bool `json:"unformatted_address"`
	MobileNumber string `json:"mobile_number"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	Url string `json:"url"`
	Latitude string `json:"latitude"`
	Longtitude string `json:"longtitude"`
	EmergencyContactName string `json:"emergency_contact_name"`
	EmergencyContactPhoneNumber string `json:"emergency_contact_phone_number"`
	EmergencyContactRelationship string `json:"emergency_contact_relationship"`
	OnWhatsapp bool `json:"on_whatsapp"`
  }