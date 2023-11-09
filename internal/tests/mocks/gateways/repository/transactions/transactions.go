package transactionsmocks

import (
	"dmorsoleto/internal/entity"

	"github.com/stretchr/testify/mock"
)

type TransactionsRepositoryMock struct {
	mock.Mock
}

func (ref *TransactionsRepositoryMock) Add(account entity.Transactions) (string, error) {
	args := ref.Called(account)
	return args.Get(0).(string), args.Error(1)
}
