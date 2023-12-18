package models

import "time"

type Currency struct {
	Id string `json:"id"`
	Status int64 `json:"status"`
	Sort int64 `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	CreatorId string `json:"creator_id"`
	ModifierId string `json:"modifier_id"`
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	StandardPrecision int64 `json:"standard_precision"`
	PricePrecision int64 `json:"price_precision"`
	Position string `json:"position"`
  }