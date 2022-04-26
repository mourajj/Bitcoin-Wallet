package models

type Funds struct {
	ID       uint64  `json:"id,omitempty"`
	User     string  `json:"user,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
}
