package handlers

import (
	"bytes"
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/gateways/repository/accounts"
	accountsmocks "dmorsoleto/internal/tests/mocks/usecase/accounts"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

const (
	id = "cb1fd9fa-6d15-4a5d-8543-ae1ac1fbd563"
)

var (
	errFoo = errors.New("foo")
)

type AccountsHandlerTestSuite struct {
	suite.Suite
	accountsHandler AccountsHandler
	accountsUseCase *accountsmocks.AccountsUseCaseMock
}

func (ref *AccountsHandlerTestSuite) SetupTest() {
	ref.accountsUseCase = new(accountsmocks.AccountsUseCaseMock)
	ref.accountsHandler = NewAccountsHandler(ref.accountsUseCase)
}

func TestAccountsHandlerTestSuite(t *testing.T) {
	suiteTest := &AccountsHandlerTestSuite{}
	suite.Run(t, suiteTest)
}

func (ref *AccountsHandlerTestSuite) TestGetAccount_Sucess() {
	req, err := http.NewRequest("GET", "/accounts/123", nil)
	ref.NoError(err)

	vars := map[string]string{"accountId": "123"}
	req = mux.SetURLVars(req, vars)

	w := httptest.NewRecorder()

	expectedAccount := newAccountEntity()

	ref.accountsUseCase.On("Get", "123").Return(expectedAccount, nil)

	ref.accountsHandler.GetAccount(w, req)

	if w.Code != http.StatusOK {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expected := `{"account_id":"cb1fd9fa-6d15-4a5d-8543-ae1ac1fbd563","document_number":"123456789"}`
	if w.Body.String() != expected {
		ref.T().Errorf("Expected body %s, got %s", expected, w.Body.String())
	}

	ref.accountsUseCase.AssertExpectations(ref.T())
}

func (ref *AccountsHandlerTestSuite) TestGetAccount_UseCase_Error() {
	req, err := http.NewRequest("GET", "/account/123", nil)
	ref.NoError(err)

	vars := map[string]string{"accountId": "123"}
	req = mux.SetURLVars(req, vars)

	w := httptest.NewRecorder()

	ref.accountsUseCase.On("Get", "123").Return(entity.Account{}, errFoo)

	ref.accountsHandler.GetAccount(w, req)

	if w.Code != http.StatusInternalServerError {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	ref.accountsUseCase.AssertExpectations(ref.T())
}

func (ref *AccountsHandlerTestSuite) TestAddAccount_Success() {
	newAccount := newAddAccount()

	payloadAccount, err := json.Marshal(newAccount)
	if err != nil {
		ref.T().Fatal(err)
	}

	req, err := http.NewRequest("POST", "/account", bytes.NewBuffer(payloadAccount))
	if err != nil {
		ref.T().Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ref.accountsUseCase.On("Add", newAccount).Return(id, nil)

	ref.accountsHandler.AddAccount(w, req)

	if w.Code != http.StatusCreated {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	ref.accountsUseCase.AssertExpectations(ref.T())
}

func (ref *AccountsHandlerTestSuite) TestAddAccount_UseCase_Error() {
	newAccount := newAddAccount()

	payloadAccount, err := json.Marshal(newAccount)
	if err != nil {
		ref.T().Fatal(err)
	}

	req, err := http.NewRequest("POST", "/account", bytes.NewBuffer(payloadAccount))
	if err != nil {
		ref.T().Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ref.accountsUseCase.On("Add", newAccount).Return(id, errFoo)

	ref.accountsHandler.AddAccount(w, req)

	if w.Code != http.StatusInternalServerError {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	ref.accountsUseCase.AssertExpectations(ref.T())
}

func (ref *AccountsHandlerTestSuite) TestAddAccount_Request_Error() {
	teste := []byte{1, 2, 3}

	req, err := http.NewRequest("POST", "/account", bytes.NewBuffer(teste))
	if err != nil {
		ref.T().Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ref.accountsHandler.AddAccount(w, req)

	if w.Code != http.StatusBadRequest {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	ref.accountsUseCase.AssertExpectations(ref.T())
}

func newAccountEntity() entity.Account {
	return entity.Account{
		AccountID:      "cb1fd9fa-6d15-4a5d-8543-ae1ac1fbd563",
		DocumentNumber: "123456789",
	}
}

func newAddAccount() accounts.AddAccount {
	return accounts.AddAccount{
		DocumentNumber: "123456789",
	}
}
