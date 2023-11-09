package controller

import (
	"dmorsoleto/internal/controller/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunHttp(accountsHandler handlers.AccountsHandler, transactionsHandler handlers.TransactionsHandler) {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", accountsHandler.AddAccount).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", accountsHandler.GetAccount).Methods("GET")
	router.HandleFunc("/transactions", transactionsHandler.AddTransaction).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
