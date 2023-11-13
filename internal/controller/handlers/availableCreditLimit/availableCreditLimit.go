package availablecreditlimit

import (
	"dmorsoleto/internal/controller/handlers"
	"dmorsoleto/internal/entity"
	availablecreditlimit "dmorsoleto/internal/usecase/availableCreditLimit"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AvailableCreditLimitHandler interface {
	GetCredit(w http.ResponseWriter, r *http.Request)
	AddCredit(w http.ResponseWriter, r *http.Request)
}

type availableCreditLimitHandler struct {
	availableCreditLimitUseCase availablecreditlimit.AvailableCreditLimit
}

func NewAvailableCreditLimit(availableCreditLimitUseCase availablecreditlimit.AvailableCreditLimit) AvailableCreditLimitHandler {
	return &availableCreditLimitHandler{
		availableCreditLimitUseCase: availableCreditLimitUseCase,
	}
}

func (ref *availableCreditLimitHandler) GetCredit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	creditAvailableId := vars["availableCreditId"]

	hasCredit, err := ref.availableCreditLimitUseCase.Get(creditAvailableId)
	if err != nil {
		logrus.Error("Something went wrong!", err)
		handlers.ParserHttpResponse(w, 0, http.StatusBadRequest, err.Error())
		return
	}

	handlers.ParserHttpResponse(w, 1, http.StatusCreated, "Account found with success", hasCredit)
}

func (ref *availableCreditLimitHandler) AddCredit(w http.ResponseWriter, r *http.Request) {

	var creditLimit entity.AddAvailableCreditLimit

	err := json.NewDecoder(r.Body).Decode(&creditLimit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		handlers.ParserHttpResponse(w, 0, http.StatusBadRequest, err.Error())
		return
	}

	lastId, err := ref.availableCreditLimitUseCase.Add(creditLimit)
	if err != nil {
		logrus.Error("Something went wrong!", err)
		handlers.ParserHttpResponse(w, 0, http.StatusBadRequest, err.Error())
		return
	}

	response := ResponseCreditLimit{
		Id: lastId,
	}

	handlers.ParserHttpResponse(w, 1, http.StatusCreated, "Credito saved with success", response)
}
