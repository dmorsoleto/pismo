package mocks

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"

	"github.com/stretchr/testify/mock"
)

type AccountsRepositoryMock struct {
	mock.Mock
}

func (ref *AccountsRepositoryMock) Get(id string) (entity.Account, error) {
	args := ref.Called(id)
	return args.Get(0).(entity.Account), args.Error(1)
}

func (ref *AccountsRepositoryMock) Add(account accounts.AddAccount) error {
	args := ref.Called(account)
	return args.Error(0)
}
