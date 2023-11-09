package transactionsusecasemocks

import (
	"dmorsoleto/internal/entity"

	"github.com/stretchr/testify/mock"
)

type TransactionsUseCaseMock struct {
	mock.Mock
}

func (m *TransactionsUseCaseMock) Add(transaction entity.Transactions) (string, error) {
	args := m.Called(transaction)
	return args.Get(0).(string), args.Error(1)
}
