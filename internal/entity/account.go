package entity

type Account struct {
	AccountID      string `json:"account_id" db:"account_id"`
	DocumentNumber string `json:"document_number" db:"document_number"`
}
