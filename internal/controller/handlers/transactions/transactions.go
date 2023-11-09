package handlers

import (
	"dmorsoleto/internal/entity"
	"encoding/json"
	"net/http"
)

type TransactionsHandler interface {
	AddTransaction(w http.ResponseWriter, r *http.Request)
}

type transactionsHandler struct{}

func NewTransactionsHandler() TransactionsHandler {
	return &transactionsHandler{}
}

func (a *transactionsHandler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction entity.Transactions

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
