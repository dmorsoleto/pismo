package operationstype

import "dmorsoleto/internal/entity"

type OperationsType interface {
	Get(operationTypeId string) (entity.OperationsType, error)
}
