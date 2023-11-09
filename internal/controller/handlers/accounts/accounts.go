package handlers

import (
	"dmorsoleto/internal/usecase/accounts"
	"encoding/json"
	"net/http"

	accounts_repo "dmorsoleto/internal/gateways/repository/accounts"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AccountsHandler interface {
	GetAccount(w http.ResponseWriter, r *http.Request)
	AddAccount(w http.ResponseWriter, r *http.Request)
}

type accountsHandler struct {
	accountsUseCase accounts.AccountsUseCase
}

func NewAccountsHandler(accountsUseCase accounts.AccountsUseCase) AccountsHandler {
	return &accountsHandler{
		accountsUseCase: accountsUseCase,
	}
}

type ResponseData struct {
	Message string `json:"message"`
}

func (ref *accountsHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["accountId"]

	account, err := ref.accountsUseCase.Get(id)
	if err != nil {
		logrus.Error("Something went wrong!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(account)
	if err != nil {
		logrus.Error("Something went wrong!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func (ref *accountsHandler) AddAccount(w http.ResponseWriter, r *http.Request) {
	var account accounts_repo.AddAccount

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idInserted, err := ref.accountsUseCase.Add(account)
	if err != nil {
		logrus.Error("Something went wrong!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ResponseAddccountData{
		Sucess:    1,
		Message:   "Account created with success",
		AccountId: idInserted,
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
