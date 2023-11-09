package transactionsusecasemocks

import (
	"dmorsoleto/internal/gateways/repository/accounts"

	"github.com/stretchr/testify/mock"
)

type TransactionsUseCaseMock struct {
	mock.Mock
}

func (m *TransactionsUseCaseMock) Add(account accounts.AddAccount) (string, error) {
	args := m.Called(account)
	return args.Get(0).(string), args.Error(1)
}
