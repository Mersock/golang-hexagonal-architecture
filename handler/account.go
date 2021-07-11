package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mersock/golang-hexagonal-architecture/errs"
	"github.com/Mersock/golang-hexagonal-architecture/service"
	"github.com/gorilla/mux"
)

type accountHandler struct {
	accService service.AccountService
}

func NewAccountHandler(accService service.AccountService) accountHandler {
	return accountHandler{accService: accService}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	CustomerID, _ := strconv.Atoi(mux.Vars(r)["CustomerID"])

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("Request body incorrect fotmat"))
		return
	}

	req := service.NewAccountReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, errs.NewValidationError("Request body incorrect fotmat"))
		return
	}

	res, err := h.accService.NewAccount(CustomerID, req)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	CustomerID, _ := strconv.Atoi(mux.Vars(r)["CustomerID"])

	res, err := h.accService.GetAccounts(CustomerID)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}
