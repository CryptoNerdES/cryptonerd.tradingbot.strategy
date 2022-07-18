package models

type Movement struct {
	Token      string `json:"token"`
	Amount     string `json:"amount"`
	EntryPrice string `json:"entry_price"`
}
