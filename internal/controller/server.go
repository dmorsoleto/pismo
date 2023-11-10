package controller

import (
	handlers_accounts "dmorsoleto/internal/controller/handlers/accounts"
	handlers_transactions "dmorsoleto/internal/controller/handlers/transactions"
	"dmorsoleto/internal/controller/middlewares"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunHttp(accountsHandler handlers_accounts.AccountsHandler, transactionsHandler handlers_transactions.TransactionsHandler) {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", middlewares.BasicAuthMiddleware(accountsHandler.AddAccount)).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", middlewares.BasicAuthMiddleware(accountsHandler.GetAccount)).Methods("GET")
	router.HandleFunc("/transactions", middlewares.BasicAuthMiddleware(transactionsHandler.AddTransaction)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
