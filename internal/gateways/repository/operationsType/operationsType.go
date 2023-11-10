package operationstype

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/helpers/database"

	"github.com/jmoiron/sqlx"
)

const (
	selectByOperationId = "SELECT * FROM pismo.operationsTypes WHERE operation_type_id = $1"
)

type operationsType struct {
	db *sqlx.DB
}

func NewOperationsTypeRepository(databaseHelper database.DatabaseHelper) OperationsType {
	db := databaseHelper.GetConnection()

	return &operationsType{
		db: db,
	}
}

func NewOperationsTypeRepositoryFromDB(db *sqlx.DB) OperationsType {
	return &operationsType{db: db}
}

func (ref *operationsType) Get(operationTypeId string) (entity.OperationsType, error) {
	operationsType := entity.OperationsType{}

	err := ref.db.Get(&operationsType, selectByOperationId, operationTypeId)
	if err != nil {
		return entity.OperationsType{}, err
	}

	return operationsType, nil
}
