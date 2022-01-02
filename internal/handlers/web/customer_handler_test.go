package web

import (
	"bytes"
	"encoding/json"
	"hex_ddd_app/internal/core/domain"
	"hex_ddd_app/internal/mocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_HTTPCustomerHandler_Get(t *testing.T) {
	// req, err := http.NewRequest("GET", "")

	w := httptest.NewRecorder()
	custService := mocks.NewCustomerServiceMock(true)
	custHandler := NewHTTPCustomerHandler(custService)

	r := mux.NewRouter()
	r.HandleFunc("/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/customers/1", nil))

	resp := w.Result()

	if resp.StatusCode != 500 {
		t.Errorf("Expected error, got ok")
	}

	custService.ErrOut = false
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("GET", "/customers/1", nil))

	resp = w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected ok, got %d", resp.StatusCode)
	}
}

func Test_HTTPCustomerHandler_Create(t *testing.T) {
	// req, err := http.NewRequest("GET", "")

	w := httptest.NewRecorder()
	custService := mocks.NewCustomerServiceMock(true)
	custHandler := NewHTTPCustomerHandler(custService)

	r := mux.NewRouter()
	r.HandleFunc("/customers", custHandler.Create).Methods(http.MethodPost)
	r.ServeHTTP(w, httptest.NewRequest("POST", "/customers", nil))

	resp := w.Result()

	if resp.StatusCode == 200 {
		t.Errorf("Expected error, got ok")
	}

	custService.ErrOut = false
	w = httptest.NewRecorder()

	jsonBody := []byte(`{"first_name": "Zavier", "last_name": "Miller"}`)

	r.ServeHTTP(w, httptest.NewRequest("POST", "/customers", bytes.NewBuffer(jsonBody)))

	resp = w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected ok, got %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var cust domain.Customer

	json.Unmarshal(body, &cust)

	if cust.FullName() != "Zavier Miller" {
		t.Errorf("Wrong customer returned. expected: %s, got: %v", jsonBody, cust)
	}
}
