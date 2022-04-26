package models

type Balance struct {
	Currency       string  `json:"currency,omitempty"`
	Amount         float64 `json:"amount,omitempty"`
	PriceInDollars float64 `json:"price_in_dollars,omitempty"`
	PriceInEuros   float64 `json:"price_in_euros,omitempty"`
	TimeOfRateUsed string  `json:"time_of_the_rate_used,omitempty"`
	TotalEuros     float64 `json:"total_euros,omitempty"`
	TotalDollars   float64 `json:"total_dollars,omitempty"`
}
