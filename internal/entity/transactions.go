package entity

import "time"

type Transactions struct {
	TransactionID   string     `json:"transaction_id" db:"transaction_id"`
	AccountID       string     `json:"account_id" db:"account_id"`
	OperationTypeID int        `json:"operation_type_id" db:"operation_type_id"`
	Amount          float64    `json:"amount" db:"amount"`
	EventDate       *time.Time `json:"event_date" db:"event_date"`
}
