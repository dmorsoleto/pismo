package transactions

import "dmorsoleto/internal/entity"

type TransactionsRepository interface {
	Add(account entity.Transactions) (string, error)
}
