package handlers

import (
	"dmorsoleto/internal/controller/handlers"
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/usecase/transactions"
	"encoding/json"
	"net/http"
)

type TransactionsHandler interface {
	AddTransaction(w http.ResponseWriter, r *http.Request)
}

type transactionsHandler struct {
	transactionUseCase transactions.TransactionsUseCase
}

func NewTransactionsHandler(transactionUseCase transactions.TransactionsUseCase) TransactionsHandler {
	return &transactionsHandler{
		transactionUseCase: transactionUseCase,
	}
}

func (ref *transactionsHandler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction entity.Transactions

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		handlers.ParserHttpResponse(w, 0, http.StatusBadRequest, err.Error())
		return
	}

	idInserted, err := ref.transactionUseCase.Add(transaction)
	if err != nil {
		handlers.ParserHttpResponse(w, 0, http.StatusBadRequest, err.Error())
		return
	}

	responseData := ResponseDataAddTransaction{
		Id: idInserted,
	}
	handlers.ParserHttpResponse(w, 0, http.StatusCreated, "Transaction created with success", responseData)
}
