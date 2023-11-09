package accounts

import (
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"
	accountmocks "dmorsoleto/internal/tests/mocks/gateways/repository/accounts"
	uuidmocks "dmorsoleto/internal/tests/mocks/helpers/uuid"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AccountsUseCaseTestSuite struct {
	suite.Suite
	accountUseCase     AccountsUseCase
	accountsRepository *accountmocks.AccountsRepositoryMock
	uuidHelper         *uuidmocks.UuidHelperMock
}

func TestAccountsUseCaseTest(t *testing.T) {
	suiteTest := &AccountsUseCaseTestSuite{}
	suite.Run(t, suiteTest)
}

func (ref *AccountsUseCaseTestSuite) SetupTest() {

	ref.accountsRepository = new(accountmocks.AccountsRepositoryMock)
	ref.uuidHelper = new(uuidmocks.UuidHelperMock)

	ref.accountUseCase = NewAccountsUseCase(ref.accountsRepository, ref.uuidHelper)
}

func (ref *AccountsUseCaseTestSuite) TestGet_Success() {
	expectedAccount := newAccountEntity()

	ref.uuidHelper.On("IsValidUUID", "123").Return(true)
	ref.accountsRepository.On("Get", "123").Return(expectedAccount, nil)

	account, err := ref.accountUseCase.Get("123")

	ref.NoError(err)
	ref.Equal(account, expectedAccount)
}

func (ref *AccountsUseCaseTestSuite) TestGet_Invalid_UUID_Error() {

	ref.uuidHelper.On("IsValidUUID", "123").Return(false)

	account, err := ref.accountUseCase.Get("123")

	expectedError := errors.New("is not valid uuid")

	ref.Error(err)
	ref.Equal(err, expectedError)
	ref.Empty(account)
}

func (ref *AccountsUseCaseTestSuite) TestAdd_Success() {
	newAccount := newAddAccount()

	ref.accountsRepository.On("Add", newAccount).Return("123", nil)

	lastId, err := ref.accountUseCase.Add(newAccount)

	ref.NoError(err)
	ref.NotEmpty(lastId)
}

func (ref *AccountsUseCaseTestSuite) TestAdd_Empty_Error() {
	newAccount := accounts.AddAccount{}

	ref.accountsRepository.On("Add", newAccount).Return("123", nil)

	lastId, err := ref.accountUseCase.Add(newAccount)

	expectedError := errors.New("account is empty")

	ref.Error(err)
	ref.Equal(err, expectedError)
	ref.Empty(lastId)
}

func newAccountEntity() entity.Account {
	return entity.Account{
		AccountID:      uuid.NewString(),
		DocumentNumber: "123456789",
	}
}

func newAddAccount() accounts.AddAccount {
	return accounts.AddAccount{
		DocumentNumber: "123456789",
	}
}
