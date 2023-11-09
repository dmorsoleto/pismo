package handlers

import (
	"dmorsoleto/internal/controller/handlers"
	"dmorsoleto/internal/entity"
	"dmorsoleto/internal/usecase/transactions"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
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

	idInserted, err := ref.transactionUseCase.Add(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := handlers.ResponseData{
		Sucess:  1,
		Message: "Transaction created with success",
		Id:      idInserted,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		logrus.Error("Something went wrong!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseJson)
}
