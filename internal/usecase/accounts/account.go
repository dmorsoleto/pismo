package accounts

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"
	"dmorsoleto/internal/helpers/uuid"
	"errors"
)

type accountsUseCase struct {
	accountsRepository accounts.AccountsRepository
	uuidHelper         uuid.UuidHelper
}

func NewAccountsUseCase(accountsRepository accounts.AccountsRepository, uuidHelper uuid.UuidHelper) AccountsUseCase {
	return &accountsUseCase{
		accountsRepository: accountsRepository,
		uuidHelper:         uuidHelper,
	}
}

func (ref *accountsUseCase) Get(id string) (entity.Account, error) {
	if !ref.uuidHelper.IsValidUUID(id) {
		return entity.Account{}, errors.New("is not valid uuid")
	}
	return ref.accountsRepository.Get(id)
}

func (ref *accountsUseCase) Add(account accounts.AddAccount) (string, error) {
	if account == (accounts.AddAccount{}) {
		return "", errors.New("account is empty")
	}
	return ref.accountsRepository.Add(account)
}
