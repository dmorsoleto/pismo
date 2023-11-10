package operationstype

import "dmorsoleto/internal/entity"

type OperationsType interface {
	Get(operationTypeId int) (entity.OperationsType, error)
}
