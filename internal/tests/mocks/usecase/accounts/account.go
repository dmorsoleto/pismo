package accountsmocks

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"

	"github.com/stretchr/testify/mock"
)

type AccountsUseCaseMock struct {
	mock.Mock
}

func (m *AccountsUseCaseMock) Get(id string) (entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Account), args.Error(1)
}

func (m *AccountsUseCaseMock) Add(account accounts.AddAccount) (string, error) {
	args := m.Called(account)
	return args.Get(0).(string), args.Error(1)
}
