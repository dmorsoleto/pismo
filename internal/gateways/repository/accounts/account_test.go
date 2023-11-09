package accounts

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
	expextedSelectAccountByID = `SELECT \* FROM accounts WHERE account_ID = \$1`
	expextedInsertAccount     = `INSERT INTO pismo\.accounts \(document_number\) VALUES \(\?\)`
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

	ref.repository = NewAccountsRepository(sqlxDB)
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

	ref.sqlMock.ExpectExec(expextedInsertAccount).
		WithArgs(newAccount.DocumentNumber).
		WillReturnResult(successResult)

	err := ref.repository.Add(newAccount)

	ref.NoError(err)
}

func (ref *AccountsRepositoryTestSuite) TestAdd_Error() {

	newAccount := newAddAccount()

	ref.sqlMock.ExpectExec(expextedInsertAccount).
		WithArgs(newAccount.DocumentNumber).
		WillReturnError(errFoo)

	err := ref.repository.Add(newAccount)

	ref.Error(err)
	ref.ErrorIs(err, errFoo)
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
