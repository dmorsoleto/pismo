package accounts

type AddAccount struct {
	DocumentNumber string `json:"document_number" db:"document_number"`
}
