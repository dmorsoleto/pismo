package accounts

import (
	"context"
	"database/sql"
	uuidmocks "dmorsoleto/internal/tests/mocks/helpers/uuid"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

const (
	expextedSelectAccountByID = `SELECT \* FROM pismo\.accounts WHERE account_ID = \$1`
	expextedInsertAccount     = `INSERT INTO pismo\.accounts \(account_ID, document_number\) VALUES \(\$1, \$2\)`
)

var (
	errFoo = errors.New("foo")

	successResult = sqlmock.NewResult(1, 1)
)

type AccountsRepositoryTestSuite struct {
	ctx context.Context
	suite.Suite
	sqlMock    sqlmock.Sqlmock
	repository AccountsRepository
	uuidHelper *uuidmocks.UuidHelperMock
}

func TestPostgresRepository(t *testing.T) {
	suiteTest := &AccountsRepositoryTestSuite{
		ctx: context.Background(),
	}
	suite.Run(t, suiteTest)
}

func (ref *AccountsRepositoryTestSuite) SetupTest() {
	var err error
	var mockDB *sql.DB

	mockDB, ref.sqlMock, err = sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	if err != nil {
		panic(err)
	}

	ref.uuidHelper = new(uuidmocks.UuidHelperMock)

	ref.repository = NewAccountsRepositoryFromDB(sqlxDB, ref.uuidHelper)
}

func (ref *AccountsRepositoryTestSuite) TestGet_Success() {

	ref.sqlMock.ExpectQuery(expextedSelectAccountByID).
		WithArgs("123").
		WillReturnRows(newAccountRow())

	account, err := ref.repository.Get("123")

	ref.NoError(err)
	ref.NotEmpty(account)
}

func (ref *AccountsRepositoryTestSuite) TestGet_Error() {

	ref.sqlMock.ExpectQuery(expextedSelectAccountByID).
		WithArgs("123").
		WillReturnError(errFoo)

	account, err := ref.repository.Get("123")

	ref.Error(err)
	ref.ErrorIs(err, errFoo)
	ref.Empty(account)
}

func (ref *AccountsRepositoryTestSuite) TestAdd_Success() {

	newAccount := newAddAccount()

	id := "1cc76376-43d5-4d26-bf15-2f2506bf5e20"

	ref.sqlMock.ExpectExec(expextedInsertAccount).
		WithArgs(id, newAccount.DocumentNumber).
		WillReturnResult(successResult)

	ref.uuidHelper.On("Generate").Return(id)

	lastId, err := ref.repository.Add(newAccount)

	ref.NoError(err)
	ref.NotEmpty(lastId)
	ref.Equal(id, lastId)
}

func (ref *AccountsRepositoryTestSuite) TestAdd_Error() {

	newAccount := newAddAccount()

	id := "1cc76376-43d5-4d26-bf15-2f2506bf5e20"

	ref.sqlMock.ExpectExec(expextedInsertAccount).
		WithArgs(id, newAccount.DocumentNumber).
		WillReturnError(errFoo)

	ref.uuidHelper.On("Generate").Return(id)

	lastId, err := ref.repository.Add(newAccount)

	ref.Error(err)
	ref.ErrorIs(err, errFoo)
	ref.Empty(lastId)
}

func newAccountRow() *sqlmock.Rows {
	accounts := []string{"account_id", "document_number"}

	return sqlmock.NewRows(accounts).AddRow("123", "123456789")
}

func newAddAccount() AddAccount {
	return AddAccount{
		DocumentNumber: "123456789",
	}
}
