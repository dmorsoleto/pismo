package handlers

import "net/http"

type TransactionsHandler interface {
	AddTransaction(w http.ResponseWriter, r *http.Request)
}

type transactionsHandler struct{}

func NewTransactionsHandler() TransactionsHandler {
	return &transactionsHandler{}
}

func (a *transactionsHandler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	print("GetAccounts")
}
