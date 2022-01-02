package main

import (
	"hex_ddd_app/internal/core/services"
	"hex_ddd_app/internal/handlers/web"
	"hex_ddd_app/internal/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := repositories.NewInMemoryDB()

	customerService := services.NewCustomerService(db)

	custHandler := web.NewHTTPCustomerHandler(customerService)

	r := mux.NewRouter()

	r.HandleFunc("/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/customers", custHandler.Create).Methods(http.MethodPost)

	log.Println("listening at port 4200...")
	panic(http.ListenAndServe(":4200", r))
}
