package transactions

import (
	"dmorsoleto/internal/entity"
	accountmocks "dmorsoleto/internal/tests/mocks/gateways/repository/accounts"
	transactionsmocks "dmorsoleto/internal/tests/mocks/gateways/repository/transactions"
	uuidmocks "dmorsoleto/internal/tests/mocks/helpers/uuid"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	accountID = "1cc76376-43d5-4d26-bf15-2f2506bf5e20"
)

var (
	errFoo = errors.New("foo")
)

type transactionsUseCaseTestSuite struct {
	suite.Suite

	transactionsUseCase   TransactionsUseCase
	transactionRepository *transactionsmocks.TransactionsRepositoryMock
	accountsRepository    *accountmocks.AccountsRepositoryMock
	uuidHelper            *uuidmocks.UuidHelperMock
}

func TestTransactionsUseCaseTest(t *testing.T) {
	suiteTest := &transactionsUseCaseTestSuite{}
	suite.Run(t, suiteTest)
}

func (ref *transactionsUseCaseTestSuite) SetupTest() {

	ref.transactionRepository = new(transactionsmocks.TransactionsRepositoryMock)
	ref.accountsRepository = new(accountmocks.AccountsRepositoryMock)
	ref.uuidHelper = new(uuidmocks.UuidHelperMock)

	ref.transactionsUseCase = NewTransactionsUseCase(ref.transactionRepository, ref.accountsRepository, ref.uuidHelper)
}

func (ref *transactionsUseCaseTestSuite) TestAdd_Success() {
	newTransaction := newTransaction()

	ref.uuidHelper.On("IsValidUUID", accountID).Return(true)
	ref.transactionRepository.On("Add", newTransaction).Return("123", nil)
	ref.accountsRepository.On("Get", accountID).Return(newAccount(), nil)

	id, err := ref.transactionsUseCase.Add(newTransaction)

	ref.NoError(err)
	ref.NotEmpty(id)
	ref.Equal("123", id)

	ref.transactionRepository.AssertExpectations(ref.T())
	ref.accountsRepository.AssertExpectations(ref.T())
}

func (ref *transactionsUseCaseTestSuite) TestAdd_Not_Valid_UUID_Error() {
	newTransaction := newTransaction()

	ref.uuidHelper.On("IsValidUUID", accountID).Return(false)

	id, err := ref.transactionsUseCase.Add(newTransaction)

	expextedError := errors.New("is not valid uuid")

	ref.Error(err)
	ref.Empty(id)
	ref.Equal(expextedError, err)

	ref.transactionRepository.AssertExpectations(ref.T())
	ref.accountsRepository.AssertExpectations(ref.T())
}

func (ref *transactionsUseCaseTestSuite) TestAdd_AccountNotFound_Error() {
	newTransaction := newTransaction()

	ref.uuidHelper.On("IsValidUUID", accountID).Return(true)
	ref.accountsRepository.On("Get", accountID).Return(entity.Account{}, errFoo)

	id, err := ref.transactionsUseCase.Add(newTransaction)

	expectedError := errors.New("account not found")

	ref.Error(err)
	ref.Empty(id)
	ref.Equal(err, expectedError)

	ref.transactionRepository.AssertExpectations(ref.T())
	ref.accountsRepository.AssertExpectations(ref.T())
}

func newTransaction() entity.Transactions {
	return entity.Transactions{
		AccountID:       accountID,
		OperationTypeID: 1,
		Amount:          123,
	}
}

func newAccount() entity.Account {
	return entity.Account{
		AccountID:      accountID,
		DocumentNumber: "123456789",
	}
}
