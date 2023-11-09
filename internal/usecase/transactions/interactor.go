package transactions

import "dmorsoleto/internal/entity"

type TransactionsUseCase interface {
	Add(transaction entity.Transactions) (string, error)
}
