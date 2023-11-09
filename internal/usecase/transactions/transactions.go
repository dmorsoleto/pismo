package transactions

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/transactions"
	"dmorsoleto/internal/helpers/uuid"
	"errors"
)

type transactionsUseCase struct {
	transactionsRepository transactions.TransactionsRepository
	uuidHelper             uuid.UuidHelper
}

func NewTransactionsUseCase(transactionsRepository transactions.TransactionsRepository, uuidHelper uuid.UuidHelper) TransactionsUseCase {
	return &transactionsUseCase{
		transactionsRepository: transactionsRepository,
		uuidHelper:             uuidHelper,
	}
}

func (ref *transactionsUseCase) Add(transaction entity.Transactions) (string, error) {
	if !ref.uuidHelper.IsValidUUID(transaction.AccountID) {
		return "", errors.New("is not valid uuid")
	}
	return ref.transactionsRepository.Add(transaction)
}
