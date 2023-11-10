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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := handlers.ResponseData{
		Success: 1,
		Message: "Transaction created with success",
	}

	idInserted, err := ref.transactionUseCase.Add(transaction)
	if err != nil {
		response.Success = 0
		response.Message = err.Error()
		responseJson, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
		return
	}

	response.Id = idInserted

	responseJson, _ := json.Marshal(response)

	w.WriteHeader(http.StatusCreated)
	w.Write(responseJson)
}
