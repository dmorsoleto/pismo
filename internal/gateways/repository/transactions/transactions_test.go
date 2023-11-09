package transactions

import (
	"database/sql"
	"dmorsoleto/internal/entity"
	uuidmocks "dmorsoleto/internal/tests/mocks/helpers/uuid"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

const (
	id = "1cc76376-43d5-4d26-bf15-2f2506bf5e20"

	expectedInsertTransaction = `INSERT INTO pismo\.transactions \(transaction_id, account_id, operations_type_id, amount, event_date\) VALUES \(\$1, \$2\)`
)

var (
	successResult = sqlmock.NewResult(1, 1)
	errFoo        = errors.New("foo")
)

type TransactionsRepositoryTestSuite struct {
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	repository TransactionsRepository
	uuidHelper *uuidmocks.UuidHelperMock
}

func TestPostgresRepository(t *testing.T) {
	suiteTest := &TransactionsRepositoryTestSuite{}
	suite.Run(t, suiteTest)
}

func (ref *TransactionsRepositoryTestSuite) SetupTest() {
	var err error
	var mockDB *sql.DB

	mockDB, ref.sqlMock, err = sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	if err != nil {
		panic(err)
	}

	ref.uuidHelper = new(uuidmocks.UuidHelperMock)

	ref.repository = NewTransactionsRepositoryFromDB(sqlxDB, ref.uuidHelper)
}

func (ref *TransactionsRepositoryTestSuite) TestAdd_Success() {
	transaction := addNewTransaction()
	ref.uuidHelper.On("Generate").Return(id, nil)

	ref.sqlMock.ExpectExec(expectedInsertTransaction).
		WithArgs(id, "1", 1, 123.45, sqlmock.AnyArg()).
		WillReturnResult(successResult)

	lastId, err := ref.repository.Add(transaction)

	ref.NoError(err)
	ref.Equal(id, lastId)
}

func (ref *TransactionsRepositoryTestSuite) TestAdd_Error() {
	transaction := addNewTransaction()
	ref.uuidHelper.On("Generate").Return(id, nil)

	ref.sqlMock.ExpectExec(expectedInsertTransaction).
		WithArgs(id, "1", 1, 123.45, sqlmock.AnyArg()).
		WillReturnError(errFoo)

	lastId, err := ref.repository.Add(transaction)

	ref.Error(err)
	ref.Equal(err, errFoo)
	ref.Empty(lastId)
}

func addNewTransaction() entity.Transactions {
	return entity.Transactions{
		TransactionID:   id,
		AccountID:       "1",
		OperationTypeID: 1,
		Amount:          123.45,
	}
}
