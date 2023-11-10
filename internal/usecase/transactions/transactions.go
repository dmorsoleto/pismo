package transactions

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"
	operationstype "dmorsoleto/internal/gateways/repository/operationsType"
	"dmorsoleto/internal/gateways/repository/transactions"
	"dmorsoleto/internal/helpers/uuid"
	"errors"
)

type transactionsUseCase struct {
	transactionsRepository   transactions.TransactionsRepository
	accountsRepository       accounts.AccountsRepository
	operationsTypeRepository operationstype.OperationsType
	uuidHelper               uuid.UuidHelper
}

func NewTransactionsUseCase(transactionsRepository transactions.TransactionsRepository, accountsRepository accounts.AccountsRepository, operationsTypeRepository operationstype.OperationsType, uuidHelper uuid.UuidHelper) TransactionsUseCase {
	return &transactionsUseCase{
		transactionsRepository:   transactionsRepository,
		accountsRepository:       accountsRepository,
		operationsTypeRepository: operationsTypeRepository,
		uuidHelper:               uuidHelper,
	}
}

func (ref *transactionsUseCase) Add(transaction entity.Transactions) (string, error) {
	if !ref.uuidHelper.IsValidUUID(transaction.AccountID) {
		return "", errors.New("is not valid uuid")
	}

	hasAccount, err := ref.accountsRepository.Get(transaction.AccountID)
	if err != nil || hasAccount == (entity.Account{}) {
		return "", errors.New("account not found")
	}

	hasOperationType, err := ref.operationsTypeRepository.Get(transaction.OperationTypeID)
	if err != nil || hasOperationType == (entity.OperationsType{}) {
		return "", errors.New("operation type not found")
	}

	return ref.transactionsRepository.Add(transaction)
}
