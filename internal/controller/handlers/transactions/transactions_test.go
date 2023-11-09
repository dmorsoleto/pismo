package handlers

import (
	"bytes"
	"dmorsoleto/internal/entity"
	transactionsusecasemocks "dmorsoleto/internal/tests/mocks/usecase/transsactions"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	id        = "cb1fd9fa-6d15-4a5d-8543-ae1ac1fbd563"
	accountID = "1cc76376-43d5-4d26-bf15-2f2506bf5e20"
)

var (
	errFoo = errors.New("foo")
)

type TransactionsHandlerTestSuite struct {
	suite.Suite
	transactionsHandler TransactionsHandler
	transactionsUseCase *transactionsusecasemocks.TransactionsUseCaseMock
}

func (ref *TransactionsHandlerTestSuite) SetupTest() {
	ref.transactionsUseCase = new(transactionsusecasemocks.TransactionsUseCaseMock)
	ref.transactionsHandler = NewTransactionsHandler(ref.transactionsUseCase)
}

func TestTransactionsHandlerTestSuite(t *testing.T) {
	suiteTest := &TransactionsHandlerTestSuite{}
	suite.Run(t, suiteTest)
}

func (ref *TransactionsHandlerTestSuite) TestAddAccount_Success() {
	newTransaction := newTransaction()

	payloadTransaction, err := json.Marshal(newTransaction)
	if err != nil {
		ref.T().Fatal(err)
	}

	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(payloadTransaction))
	if err != nil {
		ref.T().Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ref.transactionsUseCase.On("Add", newTransaction).Return(id, nil)

	ref.transactionsHandler.AddTransaction(w, req)

	if w.Code != http.StatusCreated {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	ref.transactionsUseCase.AssertExpectations(ref.T())
}

func (ref *TransactionsHandlerTestSuite) TestAddAccount_UseCase_Error() {
	newTransaction := newTransaction()

	payloadTransaction, err := json.Marshal(newTransaction)
	if err != nil {
		ref.T().Fatal(err)
	}

	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(payloadTransaction))
	if err != nil {
		ref.T().Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ref.transactionsUseCase.On("Add", newTransaction).Return("", errFoo)

	ref.transactionsHandler.AddTransaction(w, req)

	if w.Code != http.StatusInternalServerError {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	ref.transactionsUseCase.AssertExpectations(ref.T())
}

func (ref *TransactionsHandlerTestSuite) TestAddAccount_Request_Error() {
	teste := []byte{1, 2, 3}

	req, err := http.NewRequest("POST", "/account", bytes.NewBuffer(teste))
	if err != nil {
		ref.T().Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ref.transactionsHandler.AddTransaction(w, req)

	if w.Code != http.StatusBadRequest {
		ref.T().Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	ref.transactionsUseCase.AssertExpectations(ref.T())
}

func newTransaction() entity.Transactions {
	return entity.Transactions{
		AccountID:       accountID,
		OperationTypeID: 1,
		Amount:          123,
	}
}
