package entity

type AvailableCreditLimit struct {
	AvailableCreditLimitID string  `json:"available_credit_limit_id" db:"available_credit_limit_id"`
	AccountId              string  `json:"account_id" db:"account_id"`
	AvailableCreditLimit   float64 `json:"available_credit_limit" db:"available_credit_limit"`
}

type AddAvailableCreditLimit struct {
	AccountId            string  `json:"account_id" db:"account_id"`
	AvailableCreditLimit float64 `json:"available_credit_limit" db:"available_credit_limit"`
}
