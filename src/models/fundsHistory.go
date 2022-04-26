package models

type FundsHistory struct {
	ID              uint64  `json:"id,omitempty"`
	User            string  `json:"user,omitempty"`
	Currency        string  `json:"currency,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	Date            string  `json:"date,omitempty"`
}
