package operationstypemocks

import (
	"dmorsoleto/internal/entity"

	"github.com/stretchr/testify/mock"
)

type OperationsTypeRepositoryMock struct {
	mock.Mock
}

func (ref *OperationsTypeRepositoryMock) Get(operationTypeId int) (entity.OperationsType, error) {
	args := ref.Called(operationTypeId)
	return args.Get(0).(entity.OperationsType), args.Error(1)
}
