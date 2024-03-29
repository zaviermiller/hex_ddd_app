package main

import (
	"hex_ddd_app/internal/client"
	"hex_ddd_app/internal/handlers/rest"
	"hex_ddd_app/internal/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := repositories.NewInMemoryDB()

	customerService := client.NewCustomerService(db)

	custHandler := rest.NewHTTPCustomerHandler(customerService)

	r := mux.NewRouter()

	r.HandleFunc("/customers/{id}", custHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/customers", custHandler.Create).Methods(http.MethodPost)

	log.Println("listening at port 4200...")
	panic(http.ListenAndServe(":4200", r))
}
