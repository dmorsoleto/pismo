package operationstype

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

const (
	expextedSelectByOperationId = `SELECT \* FROM pismo\.operationsTypes WHERE operation_type_id = \$1`
)

var (
	errFoo = errors.New("foo")

	successResult = sqlmock.NewResult(1, 1)
)

type OperationsTypeRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	repository OperationsType
}

func TestPostgresRepository(t *testing.T) {
	suiteTest := &OperationsTypeRepositoryTestSuite{
		ctx: context.Background(),
	}
	suite.Run(t, suiteTest)
}

func (ref *OperationsTypeRepositoryTestSuite) SetupTest() {
	var err error
	var mockDB *sql.DB

	mockDB, ref.sqlMock, err = sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	if err != nil {
		panic(err)
	}

	ref.repository = NewOperationsTypeRepositoryFromDB(sqlxDB)
}

func (ref *OperationsTypeRepositoryTestSuite) TestGet_Success() {

	ref.sqlMock.ExpectQuery(expextedSelectByOperationId).
		WithArgs("1").
		WillReturnRows(newOperatioTypeRow())

	account, err := ref.repository.Get("1")

	ref.NoError(err)
	ref.NotEmpty(account)
}

func (ref *OperationsTypeRepositoryTestSuite) TestGet_Error() {

	ref.sqlMock.ExpectQuery(expextedSelectByOperationId).
		WithArgs("1").
		WillReturnError(errFoo)

	account, err := ref.repository.Get("1")

	ref.Error(err)
	ref.ErrorIs(err, errFoo)
	ref.Empty(account)
}

func newOperatioTypeRow() *sqlmock.Rows {
	operationsType := []string{"operation_type_id", "description"}

	return sqlmock.NewRows(operationsType).AddRow("1", "PAYMENT")
}
