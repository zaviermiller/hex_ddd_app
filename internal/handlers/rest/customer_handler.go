package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"hex_ddd_app/internal/client/entities"
	"hex_ddd_app/internal/client/ports"

	"github.com/gorilla/mux"
)

// adapter for customer driver port
type HTTPCustomerHandler struct {
	customerService ports.CustomerService
}

func NewHTTPCustomerHandler(customerService ports.CustomerService) *HTTPCustomerHandler {
	return &HTTPCustomerHandler{
		customerService: customerService,
	}
}

func (h HTTPCustomerHandler) Get(w http.ResponseWriter, r *http.Request) {
	// middlewares here?

	w.Header().Set("Content-Type", "application/json")

	pathVars := mux.Vars(r)
	id := pathVars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	customer, err := h.customerService.Get(uint(idInt))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (h HTTPCustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	// middlewares here?

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if len(body) < 1 {
		http.Error(w, "req body must have customer", 400)
	}

	var cust entities.Customer
	err = json.Unmarshal(body, &cust)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	err = h.customerService.Create(&cust)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(cust)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
