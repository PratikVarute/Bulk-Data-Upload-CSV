package model

type Subcription struct {
	Plan           string `json:"plan"`
	Status         string `json:"status"`
	Payment_method string `json:"payment_method"`
	Term           string `json:"term"`
}
