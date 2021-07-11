package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mersock/golang-hexagonal-architecture/service"
	"github.com/gorilla/mux"
)

type customerHandler struct {
	custService service.CustomerService
}

func NewCustomerHandler(custService service.CustomerService) customerHandler {
	return customerHandler{
		custService: custService,
	}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custService.GetCustomers()
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	CustomerID, _ := strconv.Atoi(mux.Vars(r)["CustomerID"])
	customer, err := h.custService.GetCustomer(CustomerID)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
