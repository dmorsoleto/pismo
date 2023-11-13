package controller

import (
	handlers_accounts "dmorsoleto/internal/controller/handlers/accounts"
	handlers_credit "dmorsoleto/internal/controller/handlers/availableCreditLimit"
	handlers_transactions "dmorsoleto/internal/controller/handlers/transactions"
	"dmorsoleto/internal/controller/middlewares"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunHttp(accountsHandler handlers_accounts.AccountsHandler, transactionsHandler handlers_transactions.TransactionsHandler, availableCreditLimitHandler handlers_credit.AvailableCreditLimitHandler) {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", middlewares.BasicAuthMiddleware(accountsHandler.AddAccount)).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", middlewares.BasicAuthMiddleware(accountsHandler.GetAccount)).Methods("GET")
	router.HandleFunc("/transactions", middlewares.BasicAuthMiddleware(transactionsHandler.AddTransaction)).Methods("POST")
	router.HandleFunc("/credit/{availableCreditId}", middlewares.BasicAuthMiddleware(availableCreditLimitHandler.GetCredit)).Methods("GET")
	router.HandleFunc("/credit", middlewares.BasicAuthMiddleware(availableCreditLimitHandler.AddCredit)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
